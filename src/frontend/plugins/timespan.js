import moment from 'moment';
var timespan = {
  IsOpen(schedule, day){
    var rules = this.GetRules(schedule);
    var startDay = this.getOnlyDay(day);
    var endDay = this.getNextDate(startDay);
    var events = this.GetEventsInDay(rules, startDay, day, endDay);
    return events.length>0;
  },


  StartInDay(rules, day) {
    var startDay = this.getOnlyDay(day);
    var passed = true;
    var wd = moment(day[0]+"-"+day[1]+"-"+day[2], "YYYY-M-D").locale('en').weekday();
    for (var i = 0; i < rules.length; i++) {
      let ruleObject = rules[i];
      if(typeof ruleObject.wd !="undefined" && ruleObject.wd.length == 0 && ruleObject.wd.length == 0 && ruleObject.Timeline.length== 0) {
        return true;
      }
      passed = true;
      if(typeof ruleObject.wd != "undefined"){
        if(Array.isArray(ruleObject.wd[0])) {
          ruleObject.wd.forEach((val, index) => {
            passed &= val.indexOf(wd) != -1;
          });
        } else {
          passed &= ruleObject.wd.indexOf(wd) != -1;
        }
      }
      if(passed && typeof ruleObject.timeline != "undefined") {
        let bhr = Array.isArray(ruleObject.bhr[0])? ruleObject.bhr : [ruleObject.bhr];
        let hr = Array.isArray(ruleObject.hr[0])? ruleObject.hr : [ruleObject.hr];
        let orPass = false;
        for(let i = 0; i < ruleObject.timeline.length; i++ ){
          for(let b = 0; b < bhr.length; b++ ){
            open = this.addTime(this.deepClone(startDay), bhr[b]);

            close = this.addTime(this.deepClone(open), hr[b]);

            orPass |= !this.isLaterThan(ruleObject.timeline[i].b, open) && !this.isLaterThan(close, ruleObject.timeline[i].e)
          }
        }
        passed = passed && (orPass || ruleObject.timeline.length == 0);
      }
      if(passed) {
        return true;
      }
    };
    return false;
  },

  EventsInRange(schedule, start, end){
    var rules = this.GetRules(schedule);   
    var events = [];
    var startDay = this.getOnlyDay(start);
    var endDay = this.getOnlyDay(end);
    var finished = false;
    for (var day = startDay; !finished; day = this.getNextDate(day)) {
      events = events.concat(this.GetEventsInDay(rules, day, start, end));
      finished = JSON.stringify(day) == JSON.stringify(endDay);
    }
    return events;
  },

  addTime(date, time){
    for (let z = 3 - time.length; z > 0; z--) {
      time.push(0);
    }
    for (let z = 6 - date.length; z > 0; z--) {
      date.push(0);
    }
    var max = [23, 59, 59];
    var addDays = 0;
    var over = false;
    for (var i=2; i >= 0 ; i--) {
      var p = date[i+3] + (over? 1:0);
      if(p + time[i] > max[i]){
        if(i == 0){
          addDays = Math.floor((p + time[i]) / 24);
        }
        date[i+3] = p + time[i] - (max[i]+1)*addDays;
        over = true;
      } else {
        date[i+3] = p + time[i];
        over = false;
      }
    }
    if(over){
      var next = date;
      for (i=0; i < addDays; i++) {
        next = this.getNextDate(next);
      }
      date = [next[0],next[1],next[2],date[3],date[4],date[5]];
    }

    return date;
  },

  createDatetime(date, time){
    var ret = [0,0,0,0,0,0];
    for (var i=0; i < 3; i++) {
      if(date.length > i)
        ret[i] = date[i];
      if(time.length > i)
        ret[i+3] = time[i];
    }
    return ret;
  },

  GetEventsInDay(rules, day, start, end){
    var startDay = this.getOnlyDay(day);
    var passed = true;
    var events = [];

    var wd = moment(day[0]+"-"+day[1]+"-"+day[2], "YYYY-M-D").locale('en').weekday();
    for (var i = 0; i < rules.length; i++) {
      let open = null, close = null;
      let ruleObject = rules[i];
      if(typeof ruleObject.wd !="undefined" && ruleObject.wd.length == 0 && ruleObject.wd.length == 0 && ruleObject.Timeline.length== 0) {
        continue;
      }
      passed = true;
      if(typeof ruleObject.wd != "undefined"){
        if(Array.isArray(ruleObject.wd[0])) {
          ruleObject.wd.forEach((val, index) => {
            passed &= val.indexOf(wd) != -1;
          });
        } else {
          passed &= ruleObject.wd.indexOf(wd) != -1;
        }
      }
      let bhr = Array.isArray(ruleObject.bhr[0])? ruleObject.bhr : [ruleObject.bhr];
      let hr = Array.isArray(ruleObject.hr[0])? ruleObject.hr : [ruleObject.hr];
      if(passed && typeof ruleObject.timeline != "undefined") {
        let orPass = false;
        for(let i = 0; i < ruleObject.timeline.length; i++ ){
          for(let b = 0; b < bhr.length; b++ ){
            open = this.addTime(this.deepClone(startDay), bhr[b]);

            close = this.addTime(this.deepClone(open), hr[b]);

            orPass |= !this.isLaterThan(ruleObject.timeline[i].b, open) && !this.isLaterThan(close, ruleObject.timeline[i].e)
          }
        }
        passed = passed && (orPass || ruleObject.timeline.length == 0);
      }
      if(passed && open && close) {
        events.push({"open": open, "close": close});
      }
      if(passed && ruleObject.timeline.length == 0){
        for(let b = 0; b < bhr.length; b++ ){
          open = this.addTime(this.deepClone(startDay), bhr[b]);
          close = this.addTime(this.deepClone(open), hr[b]);
          events.push({"open": open, "close": close});
        }
      }
    };
    //console.log(startDay, events);
    return events;
  },

  getNextDate(date) {
    var days_in_month = [0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
    var result = [date[0],date[1],date[2]];
    result[2] += 1;
    var dmax = days_in_month[result[1]];
    if (result[2] > dmax) {
      if (result[1] == 2) {
        if ( ((result[0] % 4) == 0) && (((result[0] % 100) != 0) || ((result[0] % 400) == 0))) {
  dmax = 29;
}
}
if (result[2] > dmax ) {
  result[2] = 1;
  result[1] += 1;
}
if (result[1] > 12) {
  result[0] += 1;
  result[1] = 1;
}
}
return result;
},

getOnlyDay(day){
  day = day.slice(0, 3);
  for(let z = 3 - day.length; z > 0; z--) {
    day.push(1);
  }
  return day;
},

GetRules(schedule){
  let recursionFind = this.recursion(schedule, false);
  let recursionRules = ( recursionFind["rules"].length > 0 )?
  recursionFind.rules : [ recursionFind["obj"] ];
  for (let i = 0; i < recursionRules.length; i++) {
    let eopen = [];
    let eclose = [];
    let r = recursionRules[i];
    let attr = ["bev", "eev", "nbev", "neev"];
    for (let a = 0; a < attr.length; a++) {
      let at = attr[a];
      if(typeof r[at] != "undefined"){
        if(Array.isArray(r[at][0])) {
          r[at].forEach((val, index) => {
            if(a<2)              
              eopen.push({ key: at, value: val});
            else
              eclose.push({ key: at.substr(1), value: val});
          });
        } else {
          if(a<2)            
            eopen.push({ key: at, value: r[at]})
          else
            eclose.push({ key: at.substr(1), value: r[at]})
        }  
        delete recursionRules[i][at];
      }
    }
    
    eopen = eopen.sort((a,b)=>{
      return this.isLaterThan(a.value, b.value)? 1 : -1;
  });
          if(eopen.length > 0 && eopen[0].key == "eev")
            eopen.unshift({key:"bev", value: [2019]});
          
          eclose = eclose.sort((a,b)=>{
            return this.isLaterThan(a.value, b.value)? 1 : -1;
        });
      //if(eclose.length == 0)
      //  eclose.push({key:"bev", value: [2019]});
      if(eclose.length > 0 && eclose[0].key == "eev")
        eclose.unshift({key:"bev", value: [2019]});

      recursionRules[i].eopen = eopen;
      recursionRules[i].eclose = eclose;
      recursionRules[i].timeline = this.flatLine(eopen, eclose);
      //console.log(JSON.stringify(recursionRules[i].timeline));
    }
    return recursionRules;
  },

  find(where, what, from){
    if(from >= where.length || from == -1) return -1;
    for (var i = from; i < where.length; i++) {
      if(where[i].key == what)
        return i;
    }
    return -1;
  },

  uglifyTimeline(e){
    let search = "bev", bev = null, timeline = [];
    for (var i = 0; i < e.length; i++) {
      if(e[i].key == search && search == "bev"){
        search = "eev";
        bev = e[i].value;
      }
      if(e[i].key == search && search == "eev"){
        search = "bev";
        timeline.push({b: bev, e: e[i].value});
        bev = null;
      }
    }
    if(bev != null){
      timeline.push({ b: bev ,e: [moment().year()+10]});
    }
    return timeline;
  },

  flatLine(eopen, eclose){
    let ueopen = this.uglifyTimeline(eopen)   // remove duplicates
    let ueclose = this.uglifyTimeline(eclose) // remove duplicates

    if(eopen.length == 0) return ueclose;
    if(eclose.length == 0) return ueopen;

    //console.log(eopen, eclose, ueopen, ueclose);
    let timeLine = [];

    for (let i = 0; i < ueopen.length; i++) {
      let o = ueopen[i];
      let foundedIndex = this.isInEvents(o.b, ueclose);
      if (foundedIndex != -1) {
        if (this.isLaterThan(ueclose[foundedIndex].e, o.e)) {
          timeLine.push(o);
        } else {
          timeLine.push({b: o.b, e: ueclose[foundedIndex].e});
          timeLine = timeLine.concat(this.findInNot(ueclose, foundedIndex+1, o.e));
        }
      }
    }

    return timeLine;
  },

  isInEvents(val, eventsNot) {
    for (let i = 0; i < eventsNot.length; i++) {
      if (!this.isLaterThan(eventsNot[i].b, val) && !this.isLaterThan(val, eventsNot[i].e)) {
        return i;
      }
    }
    return -1;
  },


  findInNot(closed, index , open) {
    var timeLine = [];
    
    for (let i = index; i < closed.length; i++) {
      let c = closed[i];
      if (!this.isLaterThan(c.b, open)) {
        if (this.isLaterThan(c.e, open)) {
          timeLine.push({b: c.b, e: open});
          return timeLine;
        } else {
          timeLine.push({b: c.b, e: c.e});
        }
      } else {
        return timeLine;
      }
    };
    return timeLine;
  },

  removeDuplicates(t){
    if(t.length < 2)
      return t;
    else{
      var last = t[t.length-1];
      for (var i = t.length - 1; i >= 1; i--) {
        if(last.b === t[i].b && last.e === t[i].e){
          t.splice(t[i],1);
        }
      }
    }
    return t;
  },

  isLaterThan(event1, event2) {
    var resolved = false;
    var result = false;
    for(var i = 0; ((i < event1.length) || (i < event2.length)) && !resolved; i++) {
      if((i < event1.length) && (i < event2.length)) {
        if(event1[i] < event2[i]) {
          result = false;
          resolved = true;
        } else if(event1[i] > event2[i]) {
          result = true;
          resolved = true;
        }
      } else if (i < event1.length) {
        if (event1[i] > 0) {
          result = true;
          resolved = true;
        }
      } else {
        if (event2[i] > 0) {
          result = false;
          resolved = true;
        }
      }
    }
    return result;
  },



  recursion(schedule, not){
    var object = {};
    var newRules = [];
    let k, pom;
    var keys = Object.keys(schedule);
    for (var i = 0; i < keys.length; i++) {
      var value = schedule[keys[i]];
      switch (keys[i]) {
        case 'and':
        for (i=0; i < value.length; i++) {
          let rec = this.recursion(value[i], not);
          if(Object.keys(rec["obj"]).length>0){
            newRules = this.mergeRulesArrays(newRules, rec.obj);
          }
          if(Object.keys(rec["rules"]).length>0){
            let cloneRules = this.deepClone(newRules);
            for (let r=0; r < rec["rules"].length; r++) {
              let pomRules = this.deepClone(cloneRules);
              if(r==0){
                newRules = this.mergeRulesArrays(pomRules, rec["rules"][r]);
              } else{
                pom = this.mergeRulesArrays(pomRules, rec["rules"][r]);
                newRules = newRules.concat(pom);
              }
            }
          }
        }
        break;
        case 'or':
        for (i=0; i < value.length; i++) {
          let rec = this.recursion(value[i], not);
          if(Object.keys(rec["obj"]).length>0){
            newRules.push(rec["obj"]);
          }
          if(Object.keys(rec["rules"]).length>0){
            newRules = newRules.concat(rec["rules"]);
          }
        }
        break;
        case 'not':
        for (i=0; i < value.length; i++) {
          let rec = this.recursion(value[i], !not);

          if(Object.keys(rec["obj"]).length>0){
            newRules = this.mergeRulesArrays(newRules, rec["obj"]);
          }
          if(rec["rules"].length>0){
            let cloneRules = newRules;
            for (let r=0; r < Object.keys(rec["rules"]).length; r++) {
              let pomRules = cloneRules;
              if(r==0){
                newRules = this.mergeRulesArrays(pomRules, rec["rules"][r]);
              } else{
                pom = this.mergeRulesArrays(pomRules, rec["rules"][r]);
                newRules = newRules.concat(pom);
              }
            }
          }
        }
        break;
        case 'eev':
        k = not? "nbev" : keys[i];
        object[k] = value;
        break;
        case 'bev':
        k = not? "neev" : keys[i];
        object[k] = value;
        break;
        case 'wd':
        if(not){
          pom = [1,2,3,4,5,6,7];
          value.forEach((wd, index) => {
            delete(pom[wd]);
          });
        }
        object[keys[i]] = not? pom : value;
        break;
        case 'hr':
        case 'bhr':
        object[keys[i]] = value;
        break;
      }
    }
    return {"obj": object, "rules": newRules};
  },

  mergeRulesArrays(rules, object){
    var pom = this.deepClone(rules);
    if(rules.length == 0){
      pom.push({});
    }
    for (let i=0; i < pom.length; i++) {
      pom[i] = this.myArrayMerge(pom[i], object);
    }
    return pom;
  },

  myArrayMerge(object1, object2){
    var keys = Object.keys(object2);
    keys.forEach((key, index) => {
      if(typeof object1[key] != "undefined"){
        if(Array.isArray( object1[key][0]) ){
          object1[key].push(object2[key]);
        } else {
          object1[key] = [object1[key], object2[key]];
        }
      } else {
        object1[key] = object2[key];
      }
    });
    return object1;
  },

  deepClone(array){
    return JSON.parse(JSON.stringify(array));
  }

}
export default timespan;
