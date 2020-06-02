/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*
* Based on: fragments of FIWARE POI Data Provider, License Apache 2,
* https://fiware-poidataprovider.readthedocs.io/en/master/index.html
*/

package utils

import (
  "fmt"
  "os"
  "sort"
  "strconv"
  "time"

  "github.com/tidwall/gjson"
  "github.com/xeipuuv/gojsonschema"
)

const schedule_schema = `{"title": "schedule", "properties": {"or": {"description": "union - valid when any of subschedules is valid", "type": "array", "items": {"ref": "#"} }, "and": {"description": "intersection - valid when all of subschedules are valid", "type": "array", "items": {"ref": "#"} }, "not": {"description": "complement - valid when the subschedule is not valid", "ref": "#"}, "wd": {"description": "weekday: valid on listed weekdays, 1=monday,...,7=sunday", "type": "array", "items": {"type": "integer", "minimum": 1, "maximum": 7 } }, "bhr": {"description": "begin hour [hour-integer, minute-integer, second-number]. End zeros can be omitted.", "type": "array", "items": {"type": "integer"} }, "hr": {"description": "hours from bhr [hour-integer, minute-integer, second-number]. Hours can be more than 23 End zeros can be omitted.", "type": "array", "items": {"type": "integer"} }, "bev": {"description": "begin event [year, month, day, hour, minute -integers, second-number]. End zeros can be omitted.", "type": "array", "items": {"type": "integer"} }, "eev": {"description": "end event [year, month, day, hour, minute -integers, second-number]. End zeros can be omitted.", "type": "array", "items": {"type": "integer"} }, "bdate": {"description": "begin date [ month, day] until the end of the year", "type": "array", "items": {"type": "integer"} }, "edate": {"description": "end date [ month, day] from the beginning of the year", "type": "array", "items": {"type": "integer"} } } }`

/*func main() {
  start := time.Now()
  const _ = `{"bev":[2019,10,8], "wd":[4], "bhr":[12],"hr":[20]}`
  const _ = `{"or": [{"and": [{"wd": [1]}, {"or": [{"bhr": [10], "hr": [2]}, {"bhr": [12,30], "hr": [2,30]} ] } ] }, {"and": [{"wd": [2, 3, 4, 5]}, {"or": [{"bhr": [8], "hr": [4]}, {"bhr": [12,30], "hr": [2,30]} ] } ] }, {"wd": [6], "bhr": [8], "hr": [4]} ] }`
  const json = `{"and": [{"eev": [2019, 10, 14]}, {"and": [{"not": [{"eev": [2019, 10, 7], "wd": [3,4,5]} ]}, {"or": [{"wd": [1,2], "bhr":[2]}, {"wd": [6,7], "bhr":[4]} ]}, {"or": [{"wd": [1,6], "hr":[6]}, {"wd": [6,2], "hr":[8]} ]} ]} ]}`

  var events = EventsInRange(json, []int{2019, 10, 7}, []int{2019, 11, 14, 12, 1})
  elapsed := time.Since(start)

  log.Printf("Total %s", elapsed)

  for _, event := range events {
    fmt.Println(event.Open, event.Close)
  }
}*/

type Event struct {
  Open  []int
  Close []int
}

type Recursion struct {
  Object map[string]Value
  Rules  []map[string]Value
}

type Value struct {
  One  []int
  Many [][]int
}

/**
  Function IsOpen testing schedule and return  true if function found some event in day
  @var string scheduleString
  @var []int day
  @return bool
*/
func IsOpen(scheduleString string, day []int) bool {
  var rules = GetRules(scheduleString)
  var startDay = getOnlyDay(day)
  var endDay = getNextDate(startDay)
  var events = getEventsInDay(rules, startDay, day, endDay)
  return len(events) > 0
}

/**
  Function StartInDay testing schedule and return  true if function found event which start in day
  @var string scheduleString
  @var []int day
  @return bool
*/
func StartInDay(scheduleString string, day []int) bool {
  var rules = GetRules(scheduleString)
  var startDay = getOnlyDay(day)
  var passed = true
  var stringtime = (strconv.Itoa(day[0]) + "-" + addZero(day[1]) + "-" + addZero(day[2])) + "T00:00:00+00:00"
  timestamp, err := time.Parse(time.RFC3339, stringtime)
  if err != nil {
    os.Exit(1)
  }
  var wd int = int(timestamp.Weekday())
  if wd == 0 {
    wd = 7
  }
  for _, ruleObject := range rules {
    if len(ruleObject.Wd.One) == 0 && len(ruleObject.Wd.Many) == 0 && len(ruleObject.Timeline) == 0 {
      return true
    }
    passed = true
    if len(ruleObject.Wd.One) > 0 || len(ruleObject.Wd.Many) > 0 {
      if len(ruleObject.Wd.Many) > 0 {
        for _, val := range ruleObject.Wd.Many {
          passed = passed && inArray(wd, val)
        }
      } else {
        passed = passed && inArray(wd, ruleObject.Wd.One)
      }
    }
    if passed {
      bhr := append(ruleObject.Bhr.Many, ruleObject.Bhr.One)
      hr := append(ruleObject.Hr.Many, ruleObject.Hr.One)
      var orPass = false
      for i := 0; i < len(ruleObject.Timeline); i++ {
        for b := 0; b < len(bhr); b++ {
          open := addTime(cloneArray(startDay), bhr[b])

          close := addTime(cloneArray(open), hr[b])

          orPass = orPass || (!isLaterThan(ruleObject.Timeline[i].Open, open) && !isLaterThan(close, ruleObject.Timeline[i].Close))
        }
      }
      passed = passed && (orPass || len(ruleObject.Timeline) == 0)
    }
    if passed {
      return true
    }
  }
  return false
}

/**
  Function EventsInRange add to dateTime variable time
  @var string scheduleString
  @var []int start
  @var []int end
  @return []Event
*/
func EventsInRange(scheduleString string, start []int, end []int) []Event {
  var rules = GetRules(scheduleString)
  var events = []Event{}
  var startDay = getOnlyDay(start)
  var endDay = getOnlyDay(end)
  var finished = false
  for day := startDay; !finished; day = getNextDate(day) {
    events = append(events, getEventsInDay(rules, day, start, end)...)
    finished = sameDay(day, endDay)
  }
  return events
}

/**
  Function addTime add to dateTime variable time
  @var []int dateTime
  @var []int time
  @return []int
*/
func addTime(dateTime []int, time []int) []int {
  for z := 3 - len(time); z > 0; z-- {
    time = append(time, 0)
  }
  for z := 6 - len(dateTime); z > 0; z-- {
    dateTime = append(dateTime, 0)
  }
  var max = []int{23, 59, 59}
  var addDays = 0
  var over = false
  for i := 2; i >= 0; i-- {
    var p = dateTime[i+3]
    if over {
      p += 1
    }
    if p+time[i] > max[i] {
      if i == 0 {
        addDays = (p + time[i]) / 24
      }
      dateTime[i+3] = p + time[i] - (max[i]+1)*addDays
      over = true
    } else {
      dateTime[i+3] = p + time[i]
      over = false
    }
  }
  if over {
    var next = dateTime
    for i := 0; i < addDays; i++ {
      next = getNextDate(next)
    }
    dateTime = []int{next[0], next[1], next[2], dateTime[3], dateTime[4], dateTime[5]}
  }
  return dateTime
}

/**
  Function createDatetime take date and time and return datetime array
  @var []int date
  @var []int time
  @return []int
*/
func createDatetime(date []int, time []int) []int {
  var ret = []int{0, 0, 0, 0, 0, 0}
  for i := 0; i < 3; i++ {
    if len(date) > i {
      ret[i] = date[i]
    }
    if len(time) > i {
      ret[i+3] = time[i]
    }
  }
  return ret
}

/**
  Function getEventsInDay return all envents in concrete day, filtered by rules and start and end variables
  @var []map[string]Value rules
  @var []int day
  @var []int start
  @var []int end
  @return []Event
*/
func getEventsInDay(rules []Rule, day []int, start []int, end []int) []Event {
  var stringtime = (strconv.Itoa(day[0]) + "-" + addZero(day[1]) + "-" + addZero(day[2])) + "T00:00:00+00:00"
  timestamp, err := time.Parse(time.RFC3339, stringtime)
  if err != nil {
    os.Exit(1)
  }
  var wd int = int(timestamp.Weekday())
  if wd == 0 {
    wd = 7
  }
  var events = []Event{}
  for _, ruleObject := range rules {
    var open = createDatetime(day, ruleObject.Bhr.One)
    var close = addTime(cloneArray(open), ruleObject.Hr.One)
    var passed = isLaterThan(open, start) && isLaterThan(end, close)
    if !passed {
      continue
    }
    if passed && (len(ruleObject.Wd.One) > 0 || len(ruleObject.Wd.Many) > 0) {
      if len(ruleObject.Wd.Many) > 0 {
        for _, val := range ruleObject.Wd.Many {
          passed = passed && inArray(wd, val)
        }
      } else {
        passed = passed && inArray(wd, ruleObject.Wd.One)
      }
    }
    if passed {
      var orPass = false
      for i := 0; i < len(ruleObject.Timeline); i++ {
        if len(ruleObject.Timeline[i].Open) > 0 {
          orPass = orPass || (!isLaterThan(ruleObject.Timeline[i].Open, open) && !isLaterThan(close, ruleObject.Timeline[i].Close))
        } else {
          orPass = orPass || !isLaterThan(ruleObject.Timeline[i].Open, open)
        }
      }
      passed = passed && (orPass || len(ruleObject.Timeline) == 0)
    }

    if passed {
      events = append(events, Event{Open: open, Close: close})
    }
  }
  return events
}

/**
  Function getNextDate returns array define day after input date
  @var []int date
  @return []int
*/
func getNextDate(date []int) []int {
  var days_in_month = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
  var result = []int{date[0], date[1], date[2]}

  result[2] += 1
  var dmax = days_in_month[result[1]]
  if result[2] > dmax {
    if result[1] == 2 {
      if ((result[0] % 4) == 0) &&
        (((result[0] % 100) != 0) || ((result[0] % 400) == 0)) {
        dmax = 29
      }
    }
    if result[2] > dmax {
      result[2] = 1
      result[1] += 1
    }
    if result[1] > 12 {
      result[0] += 1
      result[1] = 1
    }
  }
  return result
}

/**
  Function getOnlyDay returns array define exact day, length 3
  @var []int day
  @return []int
*/
func getOnlyDay(day []int) []int {
  day = day[0:3]
  for z := 3 - len(day); z > 0; z-- {
    day = append(day, 1)
  }
  return day
}

type Rule struct {
  Bhr      Value
  Hr       Value
  Wd       Value
  Timeline []Event
}

type Estruct struct {
  key   string
  value []int
}

/**
  Function GetRules returns One of plain rules generated by schedule string
  @var string scheduleString
  @return []Rule
*/
func GetRules(scheduleString string) []Rule {
  var schedule = gjson.Parse(scheduleString)
  var recursionFind = recursion(schedule, false)

  var recursionRules []map[string]Value
  var rules = []Rule{}

  if len(recursionFind.Rules) > 0 {
    recursionRules = recursionFind.Rules
  } else {
    recursionRules = []map[string]Value{recursionFind.Object}
  }
  for i := 0; i < len(recursionRules); i++ {
    eopen, eclose := []Estruct{}, []Estruct{}
    var r = recursionRules[i]
    var rule = Rule{
      Bhr:      r["bhr"],
      Hr:       r["hr"],
      Wd:       r["wd"],
      Timeline: []Event{},
    }
    var attr = []string{"bev", "eev", "nbev", "neev"}
    for a := 0; a < len(attr); a++ {
      var at = attr[a]
      if rAt, ok := r[at]; ok {

        if len(rAt.Many) > 0 {
          for _, val := range rAt.Many {
            if a < 2 {
              eopen = append(eopen, Estruct{key: at, value: val})
            } else {
              eclose = append(eclose, Estruct{key: at[1:], value: val})
            }
          }
        } else {
          if a < 2 {
            eopen = append(eopen, Estruct{key: at, value: rAt.One})
          } else {
            eclose = append(eclose, Estruct{key: at[1:], value: rAt.One})
          }
        }
        delete(recursionRules[i], at)
      }
    }

    sort.Slice(eopen[:], func(i, j int) bool {
      return !isLaterThan(eopen[i].value, eopen[j].value)
    })

    if len(eopen) > 0 && eopen[0].key == "eev" {
      eopen = append([]Estruct{{key: "bev", value: []int{2019}}}, eopen...)
    }

    sort.Slice(eclose[:], func(i, j int) bool {
      return !isLaterThan(eclose[i].value, eclose[j].value)
    })

    if len(eclose) > 0 && eclose[0].key == "eev" {
      eclose = append([]Estruct{{key: "bev", value: []int{2019}}}, eclose...)
    }

    rule.Timeline = flatLine(eopen, eclose)
    rules = append(rules, rule)
  }
  return rules
}

func uglifyTimeline(e []Estruct) []Event {
  var search = "bev"
  var bev = []int{}
  var timeline = []Event{}
  for i := 0; i < len(e); i++ {
    if e[i].key == search && search == "bev" {
      search = "eev"
      bev = e[i].value
    }
    if e[i].key == search && search == "eev" {
      search = "bev"
      timeline = append(timeline, Event{Open: bev, Close: e[i].value})
      bev = []int{}
    }
  }
  if len(bev) > 0 {
    timeline = append(timeline, Event{Open: bev, Close: []int{time.Now().Year() + 10}})
  }
  return timeline
}

func find(where []Estruct, what string, from int) int {
  if from >= len(where) || from == -1 {
    return -1
  }
  for i := from; i < len(where); i++ {
    if where[i].key == what {
      return i
    }
  }
  return -1
}

func flatLine(eopen []Estruct, eclose []Estruct) []Event {
  ueclose := uglifyTimeline(eclose) // remove duplicates
  ueopen := uglifyTimeline(eopen)   // remove duplicates
  if len(eopen) == 0 {
    return ueclose
  }
  if len(eclose) == 0 {
    return ueopen
  }
  var timeLine = []Event{}
  for i := 0; i < len(ueopen); i++ {
    o := ueopen[i]
    foundedIndex := isInEvents(o.Open, ueclose)
    if foundedIndex != -1 {
      if isLaterThan(ueclose[foundedIndex].Close, o.Close) {
        timeLine = append(timeLine, o)
      } else {
        timeLine = append(timeLine, Event{Open: o.Open, Close: ueclose[foundedIndex].Close})
        timeLine = append(timeLine, findInNot(ueclose, foundedIndex+1, o.Close)...)
      }
    }
  }
  return timeLine
}

func findInNot(closed []Event, index int, open []int) []Event {
  var timeLine = []Event{}
  for i := index; i < len(closed); i++ {
    c := closed[i]
    if !isLaterThan(c.Open, open) {
      if isLaterThan(c.Close, open) {
        timeLine = append(timeLine, Event{Open: c.Open, Close: open})
        return timeLine
      } else {
        timeLine = append(timeLine, Event{Open: c.Open, Close: c.Close})
      }
    } else {
      return timeLine
    }
  }
  return timeLine
}

func isInEvents(val []int, eventsNot []Event) int {
  for i := 0; i < len(eventsNot); i++ {
    if !isLaterThan(eventsNot[i].Open, val) && !isLaterThan(val, eventsNot[i].Close) {
      return i
    }
  }
  return -1
}

func removeDuplicates(t []Event) []Event {
  if len(t) < 2 {
    return t
  } else {
    var last = t[len(t)-1]
    for i := len(t) - 1; i >= 1; i-- {
      if !sameIntArray(last.Open, t[i].Open) || !sameIntArray(last.Close, t[i].Close) {
        return t[:i]
      }
    }
  }
  return t
}

/**
  Function isLaterThan returns true, if event1 is later than event2 (not simultaneous)
  @var []int event1
  @var []int event2
  @return bool
*/
func isLaterThan(event1 []int, event2 []int) bool {

  var resolved bool = false
  var result bool = false
  for i := 0; ((i < len(event1)) || (i < len(event2))) && !resolved; i++ {
    if i < len(event1) && i < len(event2) {
      if event1[i] < event2[i] {
        result = false
        resolved = true
      } else if event1[i] > event2[i] {
        result = true
        resolved = true
      }
    } else if i < len(event1) {
      if event1[i] > 0 {
        result = true
        resolved = true
      }
    } else {
      if event2[i] > 0 {
        result = false
        resolved = true
      }
    }
  }
  return result
}

/**
  Function cloneRules returns new variable []map[string]Value copied by input variable
  @var []map[string]Value rules
  @return []map[string]Value
*/
func cloneRules(rules []map[string]Value) []map[string]Value {
  var newRules = []map[string]Value{}
  for _, r := range rules {
    var obj = map[string]Value{}
    for key, o := range r {
      obj[key] = o
    }
    newRules = append(newRules, obj)
  }
  return newRules
}

/**
  Function recursion walking over schedule json and create object Recursion with plain rules
  @var gjson.Result schedule
  @var bool not
  @return []map[string]Value
*/
func recursion(schedule gjson.Result, not bool) Recursion {
  var object = map[string]Value{}
  var newRules = []map[string]Value{}

  keys := make([]string, 0)
  for k, _ := range schedule.Map() {
    keys = append(keys, k)
  }

  for _, key := range keys {
    var value = schedule.Map()[key].Array()
    switch key {
    case "and":
      for i := 0; i < len(value); i++ {
        var rec = recursion(value[i], not)
        if len(rec.Object) > 0 {
          newRules = mergeRulesArrays(newRules, rec.Object)
        }
        if len(rec.Rules) > 0 {
          var clonedRules = cloneRules(newRules)

          for r := 0; r < len(rec.Rules); r++ {
            var pomRules = cloneRules(clonedRules)
            if r == 0 {
              newRules = mergeRulesArrays(clonedRules, rec.Rules[r])
            } else {
              var pom = mergeRulesArrays(pomRules, rec.Rules[r])
              newRules = append(newRules, pom...)
            }
          }

        }
      }
      break
    case "or":
      for i := 0; i < len(value); i++ {
        var rec = recursion(value[i], not)
        if len(rec.Object) > 0 {
          newRules = append(newRules, rec.Object)
        }
        if len(rec.Rules) > 0 {
          newRules = append(newRules, rec.Rules...)
        }
      }
      break
    case "not":
      for i := 0; i < len(value); i++ {
        var rec = recursion(value[i], !not)
        if len(rec.Object) > 0 {
          newRules = mergeRulesArrays(newRules, rec.Object)
        }
        if len(rec.Rules) > 0 {
          var cloneRules = newRules
          for r := 0; r < len(rec.Rules); r++ {
            var pomRules = cloneRules
            if r == 0 {
              newRules = mergeRulesArrays(pomRules, rec.Rules[r])
            } else {
              var pom = mergeRulesArrays(pomRules, rec.Rules[r])
              newRules = append(newRules, pom...)
            }
          }
        }
      }
      break
    case "eev":
      var k = key
      if not {
        k = "nbev"
      }
      object[k] = Value{One: toIntArray(value)}
      break
    case "bev":
      var k = key
      if not {
        k = "neev"
      }
      object[k] = Value{One: toIntArray(value)}
      break
    case "wd":
      if not {
        var pom = []int{1, 2, 3, 4, 5, 6, 7}
        for wd, _ := range value {
          pom = remove(pom, wd)
        }
        object[key] = Value{One: pom}
      } else {
        object[key] = Value{One: toIntArray(value)}
      }
      break
    default:
      object[key] = Value{One: toIntArray(value)}
      break
    }
  }
  return Recursion{Object: object, Rules: newRules}
}

/**
  Function remove return arrayInt without element at index
  @var []int arrayInt
  @var int index
  @return []int
*/
func remove(arrayInt []int, index int) []int {
  arrayInt[index] = arrayInt[len(arrayInt)-1]
  return arrayInt[:len(arrayInt)-1]
}

/**
  Function mergeRulesArrays merge all objects in array rules with variable object
  @var []map[string]Value rules
  @var map[string]Value object
  @return []map[string]Value
*/
func mergeRulesArrays(rules []map[string]Value, object map[string]Value) []map[string]Value {
  var pom = cloneRules(rules)
  if len(rules) == 0 {
    pom = append(pom, map[string]Value{})
  }
  for i := 0; i < len(pom); i++ {
    pom[i] = myArrayMerge(pom[i], object)
  }
  return pom
}

/**
  Function myArrayMerge merge object1 with object2. If attribute exist, function create array of values that attribute
  @var map[string]Value object1
  @var map[string]Value object2
  @return map[string]Value
*/
func myArrayMerge(object1 map[string]Value, object2 map[string]Value) map[string]Value {
  for key, _ := range object2 {
    if _, ok := object1[key]; ok {
      if len(object1[key].One) == 0 {
        object1[key] = Value{Many: append(object1[key].Many, object2[key].One)}
      } else {
        var pom = [][]int{}
        pom = append(pom, object1[key].One)
        pom = append(pom, object2[key].One)
        object1[key] = Value{Many: pom}
      }
    } else {
      object1[key] = object2[key]
    }
  }

  return object1
}

/**
  Function testSchedule return true if schedule has correct format
  @var gjson.Result schedule
  @return bool
*/
func testSchedule(schedule gjson.Result) bool {
  //@todo URL to schema should be loaded in other style
  //@response NO! Loading from file is next delay (po co?)
  schemaLoader := gojsonschema.NewStringLoader(schedule_schema)
  scheduleLoader := gojsonschema.NewStringLoader(schedule.Raw)
  result, err := gojsonschema.Validate(schemaLoader, scheduleLoader)

  if err != nil {
    panic(err.Error())
  }

  // Validate JSON Schedule with Schema
  if result.Valid() {
    return true
  } else {
    fmt.Printf("The document is not valid. see errors :\n")
    for _, desc := range result.Errors() {
      fmt.Printf("- %s\n", desc)
    }
  }
  return false
}

/**
  Function addZero return string for time format (with zero)
  @var int intArray
  @return string
*/
func addZero(intArray int) string {
  if intArray < 10 {
    return "0" + strconv.Itoa(intArray)
  }
  return strconv.Itoa(intArray)
}

/**
  Function cloneArray return copy intArray
  @var []int intArray
  @return []int
*/
func cloneArray(intArray []int) []int {
  var clone = []int{}
  for i := 0; i < len(intArray); i++ {
    clone = append(clone, intArray[i])
  }
  return clone
}

/**
  Function inArray return true if element is in int array
  @var int element
  @var []int array
  @return bool
*/
func inArray(element int, array []int) bool {
  for i := 0; i < len(array); i++ {
    if array[i] == element {
      return true
    }
  }
  return false
}

/**
  Function sameDay return true if day1 anf day2 are same
  @var []int day1
  @var []int day2
  @return bool
*/
func sameDay(day1 []int, day2 []int) bool {
  for i := 0; i < 3; i++ {
    if day1[i] != day2[i] {
      return false
    }
  }
  return true
}

func sameIntArray(a []int, b []int) bool {
  if len(a) != len(b) {
    return false
  }
  for i := 0; i < len(a); i++ {
    if a[i] != b[i] {
      return false
    }
  }
  return true
}

/**
  Function toIntArray return array of int from variable jsonArray
  @var []gjson.Result jsonArray
  @return []int
*/
func toIntArray(jsonArray []gjson.Result) []int {
  var intArray []int = []int{}
  for i := 0; i < len(jsonArray); i++ {
    intArray = append(intArray, int(jsonArray[i].Int()))
  }
  return intArray
}
