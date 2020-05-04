<template>
  <div class="container" v-loading="loading">
    <header-menu activeIndex="1" @metamaskUpdated="(s)=>metamaskState = s" @getObject="(b) => blockchainLoaded(b)"></header-menu>
    <div class="text-center">

      <h1>
        Wyszukiwarka przewozu osób<br>
      </h1>
      <small class="text-muted">Przejazdy na całym świecie: Autobusem, Pociągiem, Promem, Samolotem, Koleją, Gondolą, Taxi </small>
    </div>
    <div class="inputs-group master">
      <div class="location-group">
        <autocomplete-input placeholder="Skąd" icon="icon_from.png" v-bind:value="from" v-on:change="changeFrom" />
        <autocomplete-input placeholder="Dokąd" icon="icon_to.png" v-bind:value="to" v-on:change="changeTo" />
      </div>
      <div class="others-group">
        <el-popover
        placement="bottom"
        width="200"
        trigger="click"
        >
        Distance {{distance}}km
        <el-slider
        v-model="distance"
        show-stops
        :step="10"
        :min="10"
        :show-tooltip="false"
        :max="100"
        >
      </el-slider>
      <el-button slot="reference" class="location" icon="el-icon-location"></el-button>
    </el-popover>
    <el-date-picker
    class="date-picker"
    v-model="date"
    type="date"
    value-format="yyyy-MM-dd"
    format="yyyy-MM-dd"
    placeholder="Data">
  </el-date-picker>

  <section class="el-container search">
    <el-tooltip content="Miejsce z wózkiem" placement="bottom">
      <div class="pointer" @click="wheelImage = (wheelImage != '/icon-wheelchair.jpg'? '/icon-wheelchair.jpg':'/icon-wheelchair-2.jpg');">
        <img :src="wheelImage">
      </div>
    </el-tooltip>
    <el-button icon="el-icon-search" @click="search">Szukaj</el-button>
  </section>

</div>
</div>

<div class="info-box" v-if="false && founded.length==0">
  <b>Wskazówka:</b> Wpisz skąd, dokąd chcesz jechać, wybierz datę i kliknij w Search. Możemy mieć jeszcze mało ofert. Aby zobaczyć wszystkie oferty przewoźników nie wypełniaj formularza i kliknij w Search. 
</div>
<div class="results">
  <div v-for="(row, index) in founded">
    <div class="top">
      <el-tooltip content="Miejsce z wózkiem" v-if="index%2==0" placement="bottom">
        <img src="icon-wheelchair-2.jpg">
      </el-tooltip>
      <span class="wehicle">Typ pojazdu: {{vehicleType(row.RouteType)}}</span>
      <span class="company">Firma: <a :href="row.AgencyUrl">{{row.AgencyName}}</a></span>
      <span type="text" class="blue-light-button" @click="showMap(index)">Pokaż mapę</span>
    </div>

    <h3>
      {{row.FromLabel}} > {{row.ToLabel}}
      <span class="price">Cena: <span>{{row.Price}} {{row.CurrencyType}}</span></span>
      <br><small>{{row.RouteDesc}}</small>
    </h3>
    <sheduller-linear-date-selector :value="selectedDate" :blockchain="blockchain" :metamaskState="metamaskState" :trip="row"></sheduller-linear-date-selector>
  </div>
</div>
<el-dialog width="600px" :visible.sync="mapDialogVisible">
  <div slot="title">
    {{mapDialogTitle}}
    <el-button type="text" icon="el-icon-refresh" @click="centerMap">Wycentruj mapę</el-button>
    <br />
    {{fromLatLng.join(", ")}} &nbsp; - &nbsp; {{toLatLng.join(", ")}}
  </div>
  <client-only>
    <l-map ref="map" @ready="mapReady" :zoom=13 :center="mapPosition">
      <l-tile-layer url="https://maps.heigit.org/openmapsurfer/tiles/roads/webmercator/{z}/{x}/{y}.png"></l-tile-layer>
      <l-marker ref="fromMarker" :lat-lng="fromLatLng">
        <l-icon iconUrl="/icon_from.png" :iconSize="[50, 50]" :iconAnchor="[25, 50]"></l-icon>
      </l-marker>
      <l-marker ref="toMarker" :lat-lng="toLatLng">
        <l-icon iconUrl="/icon_to.png" :iconSize="[50, 50]" :iconAncohr="[25, 50]"></l-icon>
      </l-marker>
    </l-map>
  </client-only>
</el-dialog>

<no-metamask-modal text="Aby kupić bilet musisz być podłaczony do portfela Metamask. Zaloguj sie lub link. Płać tokenami FLT." :metamaskDialogVisible.sync="metamaskDialogVisible" />

</div>

</template>
<script>
  import HeaderMenu from "~/components/HeaderMenu.vue";
  import AutocompleteInput from "~/components/AutocompleteInput.vue";
  import ShedullerLinearDateSelector from "~/components/ShedullerLinearDateSelector.vue";
  import timespan from "~/plugins/timespan.js";
  import compress from "~/plugins/compress.js";
  import moment from 'moment';
  import Blockchain from "~/components/Blockchain.vue";
  import NoMetamaskModal from "~/components/NoMetamaskModal.vue";
  import staticData from "~/plugins/static-data";

  export default {
    components: {
      HeaderMenu, AutocompleteInput, ShedullerLinearDateSelector, Blockchain, NoMetamaskModal
    },
    data: function() {
      return {
        from: "",
        to: "",
        date: "",//moment("2019-12-16").format("YYYY-MM-DD"),
        selectedDate: "",
        distance: 10,
        isFromSet: false,
        isToSet: false,
        fromLatLng: [],
        toLatLng: [],
        founded: [],
        vehicleTypes: ["Tranwaj", "Metro", "Bus", "Prom", "Kolejka linowa", "Gondola", "Kolej linowo-terenowa"],
        mapDialogTitle: "",
        mapDialogVisible: false,
        mapFromLatLng: [],
        mapPosition: [47.413220, -1.219482],
        blockchain: {},
        loading: true,
        metamaskDialogVisible: false,
        metamaskState: {},
        wheelImage: "icon-wheelchair.jpg",
      };
    },
    watch: {
      metamaskState: function(n){
        if(!n.enabledMetamask){
          this.metamaskDialogVisible = true;
        } else {
          this.metamaskDialogVisible = false;
          this.loading = false;
        }
      }
    },
    created: function(){
      moment.locale("pl");
      this.loading = true;
      setTimeout(()=>{
        if(!this.metamaskState.enabledMetamask){
          this.metamaskDialogVisible = true;
          this.loading = false;
        }
      }, 3000);
    },
    methods: {
      vehicleType(i){
        let v = staticData.vehicleType.find((a)=>a.id == i);
        return v? v.name: "";
      },
      showMap: function(index){
        this.fromLatLng = [this.founded[index].FromLat, this.founded[index].FromLng];
        this.toLatLng = [this.founded[index].ToLat, this.founded[index].ToLng];
        this.mapDialogTitle = this.founded[index].RouteDesc;
        this.centerMap();
        this.mapDialogVisible = true;
      },
      centerMap: function(){
        if (typeof this.$refs != "undefined" && typeof this.$refs.map != "undefined")
          this.$refs.map.mapObject.fitBounds([this.fromLatLng, this.toLatLng], { padding: [50, 50] });
      },
      mapReady: function(){
        this.centerMap();
        this.$refs.fromMarker.mapObject.on("click", () => {
          if (!this.dragged)
            this.$refs.map.mapObject.flyTo({ lat: this.fromLatLng[0], lng: this.fromLatLng[1] }, 17, {
              animate: true,
              duration: 1.5
            });
        });
        this.$refs.toMarker.mapObject.on("click", (el) => {
          if (!this.dragged)
            this.$refs.map.mapObject.flyTo({ lat: this.toLatLng[0], lng: this.toLatLng[1] }, 17, {
              animate: true,
              duration: 1.5
            });
        });
      },
      changedSelectedDate(index, date){
        this.founded[index].selectedDate = date;
        this.$forceUpdate();
      },
      changeFrom: function(val) {
        if(val.label.length == 0)
          this.isFromSet = false;
        else{
          this.isFromSet = true
          this.fromLatLng = [val.lat, val.lng]
        }
      },
      changeTo: function(val) {
        if(val.label.length == 0)
          this.isToSet = false;
        else{
          this.isToSet = true
          this.toLatLng = [val.lat, val.lng]
        }
      },
      search: function(){
        this.founded = [];
        var data = {
          Radius: this.distance*1000
        }
        if(this.date && this.date != ""){
          data.Date = this.date.split("-");
        }
        if(this.isFromSet){
          data.FromLat = this.fromLatLng[0];
          data.FromLng = this.fromLatLng[1];
        }
        if(this.isToSet){
          data.ToLat = this.toLatLng[0];
          data.ToLng = this.toLatLng[1];
        }
        let url = process.env.apiUrl+"/v1/search";
        $.ajax({
          url: url,
          type: "GET",
          data: data,
        }).done((respond) => {
          this.selectedDate = this.date;
          this.founded = respond;
        });
      },
      blockchainLoaded: function(b){
        this.blockchain = b;
      }
    }
  }

</script>
<style lang="scss">
.container {
  font-family: "Open Sans";
  flex-direction: column;
  align-items: center;
  word-break: break-word;
}
.search .el-button{
  background-color: #00bff3;
  font-size: 14px;
  font-weight: bold;
  color: #fff;
}
.pointer{
  cursor: pointer;
  border: 1px solid #DCDFE6;
  padding: 2px 2px 0px 3px;
}
.results{
  max-width: 1024px;
  margin: 2em auto;
  >div{
    margin-bottom: 12px;
    border-radius: 5px;
    box-shadow: 0.5px 0.9px 5px 0 rgba(0, 0, 0, 0.1);
    border: solid 1px #ebebeb;  
    padding: 20px;

    h3{
      margin: 31px 0px;
      font-size: 24px;
      font-weight: 600;
      small{
        color: #707070;
      }
      .price{
        color: #7d7d7d;
        width: 220px;
        float: right;
        text-align: center;
        >span{
          color: #333333;
        }
      }
    }
    .top{
      font-size: 16px;
      color: #707070;
      .blue-light-button{
        background: #e1f5fe;
        border-radius: 5px;
        display: inline-block;
        padding: 1px 12px;
        font-size: 14px;
        font-weight: 600;
        color: #48484a;
      }
      .wehicle, .company{
        margin: 0px 21px 0px 0px;
        white-space: nowrap;
      }
      a{
        color: #707070;
        font-weight: 600;
      }
      img{
        vertical-align: top;
        float: right;
      }
    }
  }
}
@media (min-width: 1440px) {
  h1{
    margin-top: 80px;
  }
}
@media (max-width: 768px) {
  .results > div>h3{
    text-align: center;
    .price {
      display: block;
      width: auto;
      float: none;
    }
  }
}

.vue2leaflet-map.leaflet-container{
  padding-bottom: 62%;
}

.inputs-group {
  margin-top:1em;
  display: flex;
  justify-content: center;
  flex: 1;
  flex-wrap: wrap;  

  .date-picker{
    width: 116px;
    input{
      padding-right: 0px;
      border-radius: 0px;
    }
  }
  .location{
    border-radius: 0px;
    padding: 12px 12px;
  }
  .search{
   max-width: 146px;
 }
 .el-container {

  input,
  button {
    border-radius: 0px;
  }
  &:first-child {

    input,
    button {
      border-bottom-left-radius: 4px;
      border-top-left-radius: 4px;
    }
  }

  &:last-child {

    input,
    button {
      border-bottom-right-radius: 4px;
      border-top-right-radius: 4px;
    }
  }
}
}
.location-group, .others-group{
  display: flex;
  justify-content: center;
}
@media (min-width: 1024px){
  .location-group .el-container:last-child input{
    border-bottom-right-radius: 0px;
    border-top-right-radius: 0px;
  }
}
@media (max-width: 1024px){
  .master{
    margin: 1em 2% 0px 2%;
  }
  .others-group {
    margin-top: 4px;
    .el-button.location{
      border-bottom-left-radius: 4px;
      border-top-left-radius: 4px;
    }
  }
  .inputs-group{
    flex-direction: column;
  }
  .results > div{
    margin-left: 2%;
    margin-right: 2%;
  }
}
@media (min-width: 425px){
  .others-group{
    justify-content: center;
  }
}
@media (max-width: 425px){
  .location-group{
    flex-direction: column;
    > section:first-child{
      margin-bottom: 4px;
    }
  }
  .inputs-group .el-container input{
    border-radius: 4px;
  }
}
.el-dialog{
  max-width: 96%;
}
.container{
  height: 100%;
}
body{
  height: initial;
}
</style>
