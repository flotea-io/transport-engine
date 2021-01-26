/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package sync

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"flt/models"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/astaxie/beego"

	//"encoding/json"
	transport "flt/abi/transport"
	trip "flt/abi/trip"
	compress "flt/modules/compress"
	timespan "flt/utils/timespan"

	"github.com/astaxie/beego/orm"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq"
)

var gTransportAddress = common.HexToAddress(beego.AppConfig.String("addrTransport"))
var gTransportContractAbi abi.ABI
var gTripContractAbi abi.ABI
var gLastBlock uint64
var gClient *ethclient.Client
var filterAgencyId int

func Init() {
	fmt.Println("init")
	filterAgencyId = -1

	gClient = getClient()
}

type LogTripEvent struct {
	Trip      common.Address
	TripId    *big.Int
	EventType string
}

type LogCarrier struct {
	CompanyWallet common.Address
	Company       [32]byte
	Web           [32]byte
	Index         *big.Int
}

type LogPurchased struct {
	Trip      common.Address
	TripId    *big.Int
	Tickets   *big.Int
	BuyerAddr common.Address
	Price     *big.Int
	Time      *big.Int
}

type Tickets struct {
	Addresses []common.Address
	Times     []*big.Int
}

type Body struct {
	Jsonrpc string
	Id      int
	Result  string
}

func Main(clearData bool) {
	Init()
	if clearData {
		ClearData()
	}
	main()
}

func hexToString(s string) string {
	bs, _ := hex.DecodeString(s)
	return string(bytes.Trim(bs, "\x00"))
}

func hexToInt(s string) int64 {
	result, _ := strconv.ParseUint(s, 16, 64)
	return int64(result)
}

func SplitHeader(longString string) []string {
	splits := []string{}

	var l, r int
	for l, r = 0, 64; r < len(longString); l, r = r, r+64 {
		for !utf8.RuneStart(longString[r]) {
			r--
		}
		splits = append(splits, longString[l:r])
	}
	splits = append(splits, longString[l:])
	return splits
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "clear-data" {
		ClearData()
	}

	blockNumber, err := models.GetLastBlock()
	if err == nil {
		gLastBlock = blockNumber
		//fmt.Printf(strconv.FormatUint(blockNumber, 10))
	}
	setAbis()
	getEvents()

	go subscribeEvents()

}

func ClearData() {
	//orm.Debug = true
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("TRUNCATE gtfs_routes, gtfs_agency, gtfs_stops, gtfs_trips, gtfs_stop_times, routes_tickets RESTART IDENTITY CASCADE")
	r.Exec()
	//r = o.Raw("UPDATE blockchain SET block = 15391784 WHERE id = 1")
	r = o.Raw("UPDATE blockchain SET block = " + beego.AppConfig.String("firstBlock") + " WHERE id = 1")
	//r.Exec()
	//r = o.Raw("UPDATE blockchain SET block = 14906201 WHERE id = 1")
	r.Exec()
	fmt.Println("Data reseted")
}

func setAbis() {
	cA, err := abi.JSON(strings.NewReader(string(transport.TransportABI)))
	if err != nil {
		log.Fatal(err)
	} else {
		gTransportContractAbi = cA
	}

	cB, err := abi.JSON(strings.NewReader(string(trip.TripABI)))
	if err != nil {
		log.Fatal(err)
	} else {
		gTripContractAbi = cB
	}
}

func subscribeEvents() {
	fmt.Println("subscribeEvents")

	query := ethereum.FilterQuery{
		Addresses: []common.Address{gTransportAddress, common.HexToAddress(beego.AppConfig.String("addrCarriers"))},
	}

	logs := make(chan types.Log)

	sub, err := gClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		fmt.Println(err)
	}

	instance := getTransport()
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			log.Println("Server - Sending Ping")
			_, err := instance.TokenAddress(nil)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	for {
		select {
		case err := <-sub.Err():
			fmt.Println(err)
		case vLog := <-logs:
			parseLog(vLog)
			gLastBlock = vLog.BlockNumber
			models.UpdateLastBlock(gLastBlock)
		}
	}
	fmt.Println("end subscribeEvents")
}

func getEvents() {
	fmt.Println("getEvents")

	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(gLastBlock + 1),
		Addresses: []common.Address{gTransportAddress, common.HexToAddress(beego.AppConfig.String("addrCarriers"))},
	}

	logs, err := gClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		parseLog(vLog)
		gLastBlock = vLog.BlockNumber
	}
	fmt.Println("Last block: " + strconv.FormatUint(gLastBlock, 10))
	models.UpdateLastBlock(gLastBlock)
}

func loadRoute(tripAddress common.Address) string {
	var requestBody = []byte(`{"jsonrpc":"2.0", "method":"eth_call", "params": [{"to": "` + tripAddress.String() + `", "data": "0x370158ea"}, "latest"], "id":1}`)

	resp, err := http.Post(beego.AppConfig.String("infuraHttpsConnect"), "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var structBody Body
	err = json.Unmarshal(body, &structBody)
	return structBody.Result
}

func loadAndSaveRoute(tripAddress common.Address, tripId64 int64) {
	var message = loadRoute(tripAddress)
	var arr = SplitHeader(string([]byte(message)[2:]))
	var arrLength = len(arr)

	var schedule = ""
	var i = arrLength - 3
	for ; i > 0; i -= 2 { // Search start position of schedule
		if arr[i][:3] == "000" {
			i += 2
			break
		}
	}
	for ; i < arrLength-1; i += 2 {
		schedule += arr[i]
	}

	/*
		fmt.Println("s", arrLength, arr)
		fmt.Println("carrierAddress", "0x" + strings.Trim(arr[0], "0"))
		fmt.Println("carrierId", hexToInt(arr[1]))
		fmt.Println("from: ", hexToString(string(arr[2])), hexToString(string(arr[3])))
		fmt.Println("to: ", hexToString(string(arr[4])), hexToString(string(arr[5])))
		fmt.Println("price", hexToInt(arr[6]))
		fmt.Println("palces", hexToInt(arr[8]))
		fmt.Println("route type", hexToInt(arr[12]))
		fmt.Println("enabled", hexToInt(arr[10]) == 1)
		fmt.Println("description", hexToString(arr[arrLength-1]))
		fmt.Println(compress.FromHexString(schedule))
	*/
	var location = models.TripLocFloat{
		FromLat: hexToString(string(arr[2])), FromLng: hexToString(string(arr[3])),
		ToLat: hexToString(string(arr[4])), ToLng: hexToString(string(arr[5])),
	}

	route := models.GtfsRoutes{
		RouteId:    int(tripId64),
		AgencyId:   int(hexToInt(arr[1])),
		RouteDesc:  hexToString(arr[arrLength-1]),
		RouteType:  int(hexToInt(arr[12])),
		TripWallet: strings.ToLower(tripAddress.String()),
		Places:     int(hexToInt(arr[8])),
		Schedule:   compress.FromHexString(schedule),
		Enabled:    hexToInt(arr[10]) == 1,
	}

	if (beego.AppConfig.String("AgencyWallet") != "" && filterAgencyId == -1) || (filterAgencyId != -1 && route.AgencyId != filterAgencyId) {
		return
	}

	o := orm.NewOrm()
	if created, _, err := o.ReadOrCreate(&route, "RouteId"); err == nil {
		if !created {
			_, _ = o.Update(&route)
		}
	}

	fareId := models.AddGtfsFareAttributes(float64(hexToInt(arr[6]))/1000, int(hexToInt(arr[1])))
	models.AddGtfsFareRules(int(fareId), int(tripId64))

	// Stops
	stopFrom := models.AddGtfsStop(location.FromLat, location.FromLng)
	stopTo := models.AddGtfsStop(location.ToLat, location.ToLng)
	var inserted = fillGTFS(route, stopFrom, stopTo)
	if !inserted {
		fmt.Println(stopFrom, stopTo)
		models.RemoveStopsByIds([]int64{stopFrom, stopTo})
	}

	models.AddRouteStations(route.RouteId, location)

	//fmt.Println(route)
	fmt.Println("Inserted route " + tripAddress.String() + " at index: " + strconv.Itoa(route.RouteId))

	/*
		msg := models.Message{Message: "Updated route " + tripEvent.Trip.String() + " at index: " + strconv.Itoa(route.RouteId),
					Time: time.Now().Format("2006-01-02 15:04:05")}
				models.SendBroadcastMessage(msg)
	*/
}

func setFileterAgencyId(addr string, id int64) {
	agencyWallet := beego.AppConfig.String("AgencyWallet")

	if filterAgencyId == -1 && agencyWallet != "" && agencyWallet == addr {
		pom := int(id)
		filterAgencyId = pom
	}
}

func parseLog(vLog types.Log) {
	fmt.Printf("\nLog Block Number: %d Log Index: %d | ", vLog.BlockNumber, vLog.Index)
	switch vLog.Topics[0].Hex() {
	case crypto.Keccak256Hash([]byte("TripEvent(address,uint256,string)")).Hex():
		var tripEvent LogTripEvent
		err := gTransportContractAbi.Unpack(&tripEvent, "TripEvent", vLog.Data)
		if err != nil {
			error(err.Error())
			return
		}
		loadAndSaveRoute(tripEvent.Trip, tripEvent.TripId.Int64())
	case crypto.Keccak256Hash([]byte("NewCarrier(address,bytes32,bytes32,uint256)")).Hex():
		var newCarrierEvent LogCarrier
		err := gTransportContractAbi.Unpack(&newCarrierEvent, "NewCarrier", vLog.Data)
		if err != nil {
			error(err.Error())
			return
		}
		var company = trimHexToString(newCarrierEvent.Company[:])
		var web = trimHexToString(newCarrierEvent.Web[:])
		setFileterAgencyId(newCarrierEvent.CompanyWallet.String(), newCarrierEvent.Index.Int64())
		fmt.Println("New company: " + company + " web: " + web + " with wallet: " + newCarrierEvent.CompanyWallet.String() + " at index: " + newCarrierEvent.Index.String())
		agencyId := int(newCarrierEvent.Index.Int64())
		agency := models.GtfsAgency{
			AgencyId:     agencyId,
			AgencyName:   company,
			AgencyUrl:    web,
			AgencyWallet: strings.ToLower(newCarrierEvent.CompanyWallet.String()),
		}
		id, err := models.AddAgency(&agency)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted at " + strconv.FormatInt(id, 10))
	case crypto.Keccak256Hash([]byte("CarrierUpdated(address,bytes32,bytes32,uint256)")).Hex():
		var updateCarrierEvent LogCarrier
		err := gTransportContractAbi.Unpack(&updateCarrierEvent, "UpdateCarrier", vLog.Data)
		if err != nil {
			error(err.Error())
			return
		}

		var company = trimHexToString(updateCarrierEvent.Company[:])
		var web = trimHexToString(updateCarrierEvent.Web[:])
		var wallet = updateCarrierEvent.CompanyWallet.String()

		fmt.Println("Update company: " + company + " web: " + web + " with wallet: " + wallet + " at index: " + updateCarrierEvent.Index.String())

		agency := models.GtfsAgency{
			AgencyName:   company,
			AgencyUrl:    web,
			AgencyWallet: updateCarrierEvent.CompanyWallet.String(),
		}
		err = models.UpdateAgencyByWallet(&agency)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Updated company " + company)
	case crypto.Keccak256Hash([]byte("PurchasedTickets(address,uint256,uint256,address,uint256,uint256)")).Hex():
		var purchasedEvent LogPurchased
		err := gTransportContractAbi.Unpack(&purchasedEvent, "PurchasedTickets", vLog.Data)
		if err != nil {
			error(err.Error())
			return
		}
		fmt.Println("tu", purchasedEvent.Trip.String(), purchasedEvent.Time, purchasedEvent.BuyerAddr.String())

		instance, err := trip.NewTrip(purchasedEvent.Trip, gClient)

		tickets, err := instance.GetTickets0(nil)

		o := orm.NewOrm()
		var r orm.RawSeter
		r = o.Raw("DELETE FROM routes_tickets WHERE route_id =" + purchasedEvent.TripId.String())
		r.Exec()

		routeId, err := strconv.Atoi(purchasedEvent.TripId.String())
		for i := 0; i < len(tickets.Addresses); i++ {
			time, _ := strconv.Atoi(tickets.Times[i].String())

			models.AddRoutesTickets(&models.RoutesTickets{
				TicketId:    i,
				RouteId:     routeId,
				Time:        time,
				BuyerWallet: strings.ToLower(tickets.Addresses[i].String()),
			})
		}

		fmt.Println("PurchasedTickets")
	}
}

func fillGTFS(route models.GtfsRoutes, stopFrom int64, stopTo int64) bool {

	// remove old stops
	var inserted = false
	var wdArray = []int{}
	var calendar = models.GtfsCalendar{}
	var fromString = ""
	var toString = ""
	var rules = timespan.GetRules(route.Schedule)
	var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for _, rule := range rules {

		//fmt.Println("tu", rule)

		if len(rule.Wd.One) == 0 {
			wdArray = []int{1, 2, 3, 4, 5, 6, 7}
		} else {
			wdArray = rule.Wd.One
		}

		if len(rule.Bhr.One) > 0 && len(rule.Hr.One) > 0 {
			for _, tL := range rule.Timeline {
				o := orm.NewOrm()
				calendar = models.GtfsCalendar{}
				tripId := models.AddTrip(route.RouteId)

				fromString = getTimeString(rule.Bhr.One)
				toString = getAddedTimeString(rule.Bhr.One, rule.Hr.One)

				models.AddStopTime(stopFrom, tripId, 1, fromString, fromString)
				models.AddStopTime(stopTo, tripId, 2, toString, toString)

				calendar.ServiceId = int(tripId)

				for _, d := range wdArray {
					reflect.ValueOf(&calendar).Elem().FieldByName(days[d-1]).SetInt(1)
				}

				startDate := dateTimeFromArray(tL.Open, true)
				endDate := dateTimeFromArray(tL.Close, true)
				calendar.StartDate = startDate
				calendar.EndDate = endDate

				//fmt.Println("bb", i, bevV, eevV, calendar)

				o.Insert(&calendar)
				inserted = true
			}

		}
		//fmt.Println("d", calendar)
	}
	return inserted
}

func dateTimeFromArray(dt []int, start bool) time.Time {
	switch len(dt) {
	case 1:
		return time.Date(dt[0], time.Month(1), 1, 0, 0, 0, 0, time.UTC)
		break
	case 2:
		return time.Date(dt[0], time.Month(dt[1]), 1, 0, 0, 0, 0, time.UTC)
		break
	case 3:
		return time.Date(dt[0], time.Month(dt[1]), dt[2], 0, 0, 0, 0, time.UTC)
		break
	case 4:
		return time.Date(dt[0], time.Month(dt[1]), dt[2], dt[3], 0, 0, 0, time.UTC)
		break
	case 5:
		return time.Date(dt[0], time.Month(dt[1]), dt[2], dt[3], dt[4], 0, 0, time.UTC)
		break
	case 6:
		return time.Date(dt[0], time.Month(dt[1]), dt[2], dt[3], dt[4], dt[5], 0, time.UTC)
		break
	}
	return time.Date(2000, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
}

func getTimeString(time []int) string {
	var t = ""
	for i := 0; i < len(time); i++ {
		if time[i] < 10 {
			t += ":0" + strconv.FormatInt(int64(time[i]), 10)
		} else {
			t += ":" + strconv.FormatInt(int64(time[i]), 10)
		}
	}
	for i := 0; i < 3-len(time); i++ {
		t += ":00"
	}
	return t[1:]
}

func getAddedTimeString(time []int, hours []int) string {
	var base []int
	var add []int
	if len(time) > len(hours) {
		base = time
		add = hours
	} else {
		base = hours
		add = time
	}
	for i := len(add) - 1; i >= 0; i-- {
		base[i] += add[i]
		if i > 0 && base[i] > 59 {
			base[i] -= 60
			base[i-1]++
		}
	}
	return getTimeString(base)
}

func getClient() *ethclient.Client {
	client, err := ethclient.Dial(beego.AppConfig.String("infuraWssConnect"))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("we have a connection")

	return client
}

func getTransport() *transport.Transport {
	instance, err := transport.NewTransport(gTransportAddress, gClient)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transport is loaded")

	return instance
}

func trimHexToString(data []byte) string {
	return string(bytes.Trim(data, "\x00"))
}

func error(err string) {
	fmt.Println("!!!!!!!!!!!!!!!!!!! ERROR !!!!!!!!!!!!!!!!!!!")
	fmt.Println(err)
	fmt.Println("!!!!!!!!!!!!!!!!!!! ERROR !!!!!!!!!!!!!!!!!!!")
}
