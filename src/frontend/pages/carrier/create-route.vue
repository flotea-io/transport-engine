<template>
  <div class="create-route">
    <header-menu activeIndex="3" @metamaskUpdated="(s)=>metamaskState = s" @getObject="(b) => blockchainLoaded(b)"></header-menu>
    
  <div class="container-block content" :class="{'no-map': activeTabName == 'scheduler' || activeTabName == 'json'}">
    <div class="left-map d-none d-sm-block">
      <client-only>
        <l-map ref="map" @ready="mapReady" :zoom=13 :center="mapPosition">
          <l-tile-layer url="https://maps.heigit.org/openmapsurfer/tiles/roads/webmercator/{z}/{x}/{y}.png"></l-tile-layer>
          <l-marker ref="fromMarker" :lat-lng="fromLatLng" :draggable="true">
            <l-icon iconUrl="/icon_from.png" :iconSize="[50, 50]" :iconAnchor="[25, 50]"></l-icon>
          </l-marker>
          <l-marker ref="toMarker" :lat-lng="toLatLng" :draggable="true">
            <l-icon iconUrl="/icon_to.png" :iconSize="[50, 50]" :iconAncohr="[25, 50]"></l-icon>
          </l-marker>
        </l-map>
      </client-only>
    </div>
    <div class="right-form d-none d-sm-block" :class="{'error-schedule' : validation.schedule}">
      <el-tabs type="card" v-model="activeTabName" @tab-click="changedTab">
        <el-tab-pane label="Szczegóły trasy" name="params">
          <div class="details">
            <div class="inputs-group">
            <autocomplete-input placeholder="Skąd" icon="icon_from.png" v-bind:value="from" v-on:change="(data) => change(data, 'from')" />
              <el-tooltip content="Dostęp dla wózków inwalidzkich" placement="bottom">
              <el-select class="wheel" placeholder="Wybierz" v-model="fromType">
                  <img :src="wheelType(fromType)" slot="prefix" />
                  <el-option v-for="(item, index) in stopType" :key="index" :label="item" :value="index">
                  </el-option>
                </el-select>
              </el-tooltip>
            </div>
            <div class="inputs-group">
              <autocomplete-input placeholder="Dokąd" icon="icon_to.png" v-bind:value="to" v-on:change="(data) => change(data, 'to')" />
                <el-tooltip content="Dostęp dla wózków inwalidzkich" placement="bottom">
                <el-select class="wheel" placeholder="Wybierz" v-model="toType">
                  <img :src="wheelType(toType)" slot="prefix" />
                  <el-option v-for="(item, index) in stopType" :key="index" :label="item" :value="index">
                  </el-option>
                </el-select>
              </el-tooltip>
            </div>
                <el-select v-model="vehicleType" placeholder="Rodzaj pojazdu">
                  <el-option v-for="(item, index) in filteredVehicleType" :key="item.id" :label="item.name" :value="item.id">
                  </el-option>
                </el-select>
                <el-checkbox v-model="allServices" border>All servies</el-checkbox>
                <div class="places">
                  <label class="label">Ilość miejsc</label>
                  <el-input-number v-model="places" :min="1" />
                </div>
                <el-input type="textarea" placeholder="Opis trasy" maxlength="32" show-word-limit v-model="description" />
                <p v-if="validation.description" class="error">Please input description</p>
                <el-input v-model="price" :precision="2" placeholder="Cena biletu">
                  <template slot="append">FLT</template>
                </el-input>
                <p v-if="validation.price" class="error">Please input ticket price</p>
                <el-checkbox v-if="update" v-model="enabled" size="mini" label="Enabled" border></el-checkbox>
                <el-checkbox v-model="wheel" border>Obsługa osób na wózku inwalidzkim

</el-checkbox>
                <el-button @click="activeTabName = 'scheduler'; changedTab();">Przejdź do kalendarza tras</el-button>
              </div>
            </el-tab-pane>
            <el-tab-pane label="Kalendarz trasy (wyjazdy)" name="scheduler">
              <client-only>
                <full-calendar :config="config" @event-selected="selected" :event-sources="eventSources" @event-created="eventRecieved" ref="calendar"></full-calendar>
              </client-only>
            </el-tab-pane>
            <el-tab-pane label="JSON" name="json">
              <div class="json-textarea">
                <client-only placeholder="Codemirror Loading...">
                  <Codemirror :code.sync="jsonString" 
                  :visible="showCode"
                  @change="jsonChanged"
                  >
                </Codemirror>
              </client-only>
            </div>
          </el-tab-pane>
        </el-tabs>
        <div class="form-buttons" v-loading="metamaskState.userType==null">
          <el-button v-if="metamaskState.userType==1" @click="createTrip"  type="success">Dodaj trasę</el-button>
          <div class="info-box" v-if="metamaskState.userType==0">
            <b>Wskazówka:</b> nie jesteś przewoźnikiem
          </div>
        </div>
      </div>
      <div class="d-block d-sm-none small-device">Użyj większego urządzenia</div>
      <event-modal v-bind:data="eventData" v-bind:newEvent="newEvent" v-on:save="onSave" v-on:close="onClose" v-on:delete="onDelete" />

      <no-metamask-modal text="Aby zamieścić trasę musisz zalogować się do portfela Metamask. Zaloguj sie lub link." :metamaskDialogVisible.sync="metamaskDialogVisible" />
    </div>
  </div>
  </template>
  <script>
    import HeaderMenu from "~/components/HeaderMenu.vue";
    import Codemirror from '~/components/Codemirror.vue';
    import EventModal from '~/components/EventModal.vue';
    import NoMetamaskModal from "~/components/NoMetamaskModal.vue";
    import Trip from '~/components/Trip.vue';
    import moment from 'moment';
    import scheduler from '~/plugins/scheduler';
    import Event from '~/plugins/event';
    import AutocompleteInput from "~/components/AutocompleteInput.vue";
    import $ from "jquery";
    import timespan from '~/plugins/timespan';
    import compress from "~/plugins/compress";
    import staticData from "~/plugins/static-data";

    var beautify = require("json-beautify");
    var blockChain;

    export default {
      components: {
        HeaderMenu,
        EventModal,
        AutocompleteInput,
        Codemirror,
        NoMetamaskModal
      },
      mounted: function() {
        window.addEventListener('resize', this.onResize.bind(this))
      },
      data() {
        const $this = this;
        return {
          places: 1,
          allServices: false,
          price: "",
          update: false,
          vehicleType: 42,
          schedule: {},
          description: "",
          enabled: true,
          activeTabName: 'params',
          from: "",
          to: "",
          mapPosition: [47.413220, -1.219482],
          fromLatLng: [50.05250281819231, 19.98892664909363],
          toLatLng: [52.520856132425436, 13.406206369400024],
          events: [],
          lastEventId: 0,
          newEvent: false,
          selectedEventIndex: null,
          eventData: {
            dialogVisible: false,
            checkedWd: [],
            evRange: [],
            dateRange: [],
            hrRange: [],
            mdRange: [
            [
            [3, 15],
            [6, 7]
            ]
            ]
          },
          eventSources: [{ events(start, end, timezone, callback) { $this.calendarEvents(start, end, timezone, callback); } }],
          config: {
            timezone: 'UTC',
        // height: '400',
        locale: "pl",
        firstDay: 1,
        editable: false,
        timeFormat: 'HH(:mm)',
        axisFormat: 'HH:mm',
        selectAllow: function(info) {
          if (info.start.isBefore(moment(0, "HH")))
          return false;
          return true;          
        },
        slotLabelFormat: "HH:mm",
        slotDuration: '01:00:00',
        defaultView: "agendaWeek",
        eventBackgroundColor: "#00bff3",
        displayEventTime: false,
      },
      scheduleObject: {},
      validation: {  },
      showCode: false,
      isJsonFromInput: false,
      loading: true,
      metamaskDialogVisible: false,
      metamaskState: {},
      fromType: null,
      toType: null,
      wheel: false,
      stopType: ["Brak informacji","Dostępny","Niedostępny"],
    };
  },
  watch: {
    price: function(newVal, oldVal){
      newVal = newVal.trim().replace(",",".");
      let pos = newVal.indexOf(".");
      if(pos != -1 && newVal.length - pos > 4){
        this.price = oldVal;
        $.jGrowl("Max precizion is 3.", { header: 'Error', life: 2000 });
      } else if(isNaN(newVal)){
        this.price = oldVal;
        $.jGrowl("Please input correct number.", { header: 'Error', life: 2000 });
      }
      else{
        delete this.validation.price;
        this.price = newVal.replace(".",",");
      }
    },
    metamaskState: function(n){
        if(!n.enabledMetamask){
          this.metamaskDialogVisible = true;
        } else {
          this.metamaskDialogVisible = false;
          this.loading = false;
        }
      },
    description: function(newVal, oldVal){
      if(newVal.trim().length > 0)
        delete this.validation.description;
    },
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
  computed: {
    jsonString: function(){
      let json = {
        fromLatLng: this.fromLatLng,
        toLatLng: this.toLatLng,
        vehicleType: this.vehicleType,
        places: this.places,
        price: this.price,
        description: this.description,
        schedule: this.scheduleObject
      };
      return beautify(json, null, 2, 50);
    },
    filteredVehicleType: function(){
      let filter = [1, 114, 42, 89, 104];
      let v = staticData.vehicleType;
      v.sort((a,b)=>{
        if(a.name > b.name) return 1;
        if(a.name < b.name) return -1;
        return 0;
      });
      if(this.allServices)
        return v;
      else
        return v.filter((el) => filter.indexOf(el.id) != -1);
    }
  },
  methods: {
    wheelType(i){
      return (i == 1)? "/icon-wheelchair-2.jpg" : "/icon-wheelchair.jpg";
    },
    onResize: function() {
      $(this.$refs.calendar.$el).fullCalendar('option', 'height', document.getElementById("pane-scheduler").offsetHeight - 23);
    },
    changedTab: function() {
      if (this.activeTabName == "scheduler") {
        setTimeout(() => {
          this.onResize();
        }, 10);
      }
      this.showCode=false;
      if (this.activeTabName == "json") {
        this.showCode=true;
      }
    },
    change: function(val, type) {
      this[type + "LatLng"] = [val.lat, val.lng];
      this.centerMap();
    },
    mapReady: function(map) {
      this.centerMap();
      this.$refs.fromMarker.mapObject.on("click", () => {
        if (!this.dragged)
          this.$refs.map.mapObject.flyTo({ lat: this.fromLatLng[0], lng: this.fromLatLng[1] }, 17);
      });
      this.$refs.toMarker.mapObject.on("click", (el) => {
        if (!this.dragged)
          this.$refs.map.mapObject.flyTo({ lat: this.toLatLng[0], lng: this.toLatLng[1] }, 17);
      });

      this.$refs.fromMarker.mapObject.on('dragend', (data) => {
        this.dragged = true;
        setInterval(() => this.dragged = false, 200);
        this.fromLatLng = [data.target._latlng.lat, data.target._latlng.lng];
        this.geoReverse("from", this.fromLatLng);
      });

      this.$refs.toMarker.mapObject.on('dragend', (data) => {
        this.dragged = true;
        setInterval(() => this.dragged = false, 200);
        this.toLatLng = [data.target._latlng.lat, data.target._latlng.lng];
        this.geoReverse("to", this.toLatLng);
      });
      this.geoReverse("from", this.fromLatLng);
      this.geoReverse("to", this.toLatLng);
    },
    geoReverse(type, data) {
      let url = "https://geo.flotea.pl/v1/reverse";
      $.ajax({
        url: url,
        type: "GET",
        data: {
          "point.lat": data[0],
          "point.lon": data[1],
        },
      }).done((respond) => {
        if (respond.features.length > 0) {
          this[type] = respond.features[0].properties.label;
        }
      });
    },
    centerMap: function() {
      if (typeof this.$refs != "undefined" && typeof this.$refs.map != "undefined")
        this.$refs.map.mapObject.fitBounds([this.fromLatLng, this.toLatLng], { padding: [50, 50] });
    },
    onSave() {
      //console.log(this.eventData);
      //timespan.EventsInRange(timespan.deepClone(this.events[i].object), from, to);
      //return;
      this.newEvent = false;
      this.events[this.selectedEventIndex].updateFromModal(this.eventData);
      this.eventData.dialogVisible = false;
      this.$refs.calendar.fireMethod('refetchEvents');
    },
    onClose() {
      if (this.newEvent) {
        this.events.splice(this.selectedEventIndex, 1);
      }
      this.$refs.calendar.fireMethod('refetchEvents');
    },
    onDelete() {
      this.eventData.dialogVisible = false;
      if (this.newEvent) return;
      this.events.splice(this.selectedEventIndex, 1);
      this.$refs.calendar.fireMethod('refetchEvents');
    },
    selected(event, jsEvent, view) {
      console.log(event);
      this.selectedEventIndex = this.events.findIndex(e => e.id == event.myId);
      this.newEvent = false;
      this.showEvent(this.selectedEventIndex);
    },
    showEvent(id) {
      if(id == -1){ console.log("tady"); return;}
      if(this.isJsonFromInput){
        $.jGrowl("You can edit events only which are created by calendar.", { header: 'Error', life: 3000 })
        return;
      }
      this.eventData = this.events[id].modal;
      this.eventData.dialogVisible = true;
    },
    eventRecieved(event) {
      console.log(event);
      let openCreateDialog = (event) => {
        let newEvent = new Event(this.lastEventId++, event);
        this.selectedEventIndex = this.events.length;
        this.events.push(newEvent);

        this.newEvent = true;
        this.eventData = newEvent.modal;
        this.eventData.dialogVisible = true;
        //this.$refs.calendar.fireMethod('refetchEvents');
      };

      if(this.isJsonFromInput){
        this.$confirm('You cant edit inserted schedule, do you want remove all and start from begin?', 'Warning', {
          confirmButtonText: 'Yes',
          cancelButtonText: 'No',
          type: 'warning'
        }).then(() => {
          this.events = [];
          this.isJsonFromInput = false;
          openCreateDialog(event);
        }).catch(() => {});
      } else openCreateDialog(event);
    },
    removeZeroInArray(arr, month) {
      if (month && arr.length > 1) {
        arr[1] += 1;
      }
      for (var i = arr.length - 1; i >= 0; i--) {
        if (arr[i] != 0)
          return arr.slice(0, i + 1);
      }
    },
    calendarEvents(start, end, timezone, callback) {
      let from = start.toArray(), to = end.toArray();
      from[1]++; to[1]++;
      let schedulerEvents = [];
      this.scheduleObject = {};
      if(this.isJsonFromInput){
        let foundedEvents = timespan.EventsInRange(timespan.deepClone(this.scheduleObject), from, to);
        for (let fe = 0; fe < foundedEvents.length; fe++) {
          foundedEvents[fe].open[1]--; foundedEvents[fe].close[1]--;
          schedulerEvents.push({
            title: "Plan osobisty",
            myId: 0,
            start: moment(foundedEvents[fe].open).format("Y-MM-DD HH:mm:ss"),
          //end: moment(events[i].close).format("Y-MM-DD HH:mm:ss")
        });
        }
        callback(schedulerEvents);
        return;
      }
      if (this.events.length == 0) {
        callback([]);
        return;
      }
      
      //console.log(this.scheduleObject, this.events);
      for (var i = 0; i < this.events.length; i++) {
        if (this.events.length == 1)
          this.scheduleObject = this.events[i].object;
        else{
          if(typeof this.scheduleObject.or == "undefined")
            this.scheduleObject = {"or":[]};
          this.scheduleObject.or.push(this.events[i].object);
        }

        if(Object.keys(this.scheduleObject).length == 0) continue;
        delete this.validation.schedule;
        //console.log(JSON.stringify(this.events[i].object));
        let foundedEvents = timespan.EventsInRange(timespan.deepClone(this.events[i].object), from, to);
        //console.log(JSON.stringify(timespan.deepClone(this.events[i].object), null, 2));
        //console.log(foundedEvents);
          //console.log(foundedEvents);
        for (let fe = 0; fe < foundedEvents.length; fe++) {
          foundedEvents[fe].open[1]--; foundedEvents[fe].close[1]--;
          schedulerEvents.push({
            title: "Plan #" + (i+1),
            myId: this.events[i].id,
            start: moment(foundedEvents[fe].open).format("Y-MM-DD HH:mm:ss"),
          //end: moment(events[i].close).format("Y-MM-DD HH:mm:ss")
        });
        }
      }

    //return;
      //console.log(JSON.stringify(timespan.deepClone(this.scheduleObject), null, 2));
      callback(schedulerEvents);
    },
    blockchainLoaded: function (b) {
      blockChain = b;
      blockChain.loadContract("Transport");
    },
    createTrip: function () {
      if(this.validateForm()){
          let compressed = compress.toHexArray(timespan.deepClone(this.scheduleObject));
          let location = [
          blockChain.toHex(this.fromLatLng[0]).substr(0, 22), 
          blockChain.toHex(this.fromLatLng[1]).substr(0, 24), 
          blockChain.toHex(this.toLatLng[0]).substr(0, 22), 
          blockChain.toHex(this.toLatLng[1]).substr(0, 24)
          ];
          console.log(compressed, location, this.metamaskState.userAddress);
          blockChain.sendContract("Transport", "createTrip", {from: this.metamaskState.userAddress}, ()=> {

         }, null,
         location,
         parseFloat(this.price.replace(",",".")) * 1000, 
         compressed,
         this.places,
         blockChain.toHex(this.description), 
         this.vehicleType, 
         this.enabled
         );
      }
    },
    validateForm(){
      let val = {};
      if(this.metamaskState.userType != 1){
        $.jGrowl("You are not a carrier", { header: 'Error', life: 3000 });
        this.validation.userType = true;
      }
      if(this.description.trim().length == 0)
        val.description = true;
      else
        delete val.description;

      if(this.price.trim().length == 0)
        val.price = true;
      else
        delete val.price;

      if(Object.keys(val).length == 0 && Object.keys(this.scheduleObject).length == 0){
        this.activeTabName = "scheduler";
        $.jGrowl("Please create route schedule", { header: 'Error', life: 3000 });
        val.schedule = true;
      }
      this.validation = val;
      return Object.keys(val).length == 0;
    },
    jsonChanged: function(json){

      this.isJsonFromInput = JSON.stringify(this.scheduleObject) != JSON.stringify(json.schedule);
      this.description = typeof json.description == "undefined"? "" : json.description;
      this.vehicleType = typeof json.vehicleType == "undefined"? 3 : json.vehicleType;
      this.places = typeof json.places == "undefined"? 1 : json.places;
      this.price = typeof json.price == "undefined"? "10" : json.price;
      if(this.isJsonFromInput)
        this.scheduleObject = typeof json.schedule == "undefined"? {} : json.schedule;
      this.$refs.calendar.fireMethod('refetchEvents');
      this.fromLatLng = typeof json.fromLatLng == "undefined"? [50.05250281819231, 19.98892664909363] : json.fromLatLng;
      this.toLatLng = typeof json.toLatLng == "undefined"? [52.520856132425436, 13.406206369400024] : json.toLatLng;
      this.geoReverse("from", this.fromLatLng);
      this.geoReverse("to", this.toLatLng);
      //this.events = createEventsFromJson(timespan.deepClone(this.scheduleObject));
    }
  },

}

</script>
<style lang="scss">
@import '~/node_modules/fullcalendar/dist/fullcalendar.css';
.inputs-group{
  display: flex;
    justify-content: center;
    flex: 1;
    flex-wrap: wrap;
}
.wheel{
  width: 178px;
  div.el-input--prefix input{
    padding-left: 35px;
  }
  .el-input__prefix{
    overflow: hidden;
    img{
      margin-top: 8px;
      width: 26px;
    }
  }
}
.create-route{
  height: 100%;
}
.error{
  margin: -9px 0px 0px 0px;
  color: #F56C6C;
  font-size: 12px;
}
.error-schedule #tab-scheduler{
  color: #F56C6C;
}

.json-textarea{
  margin: 0px 23px 23px 23px;
  max-height: calc(100% - 23px);
  overflow: auto;
}
.json-textarea textarea{
  border-color: transparent;
  &:focus, &:hover{
    border-color: transparent!important;
  }
}

.el-tabs.el-tabs--card.el-tabs--top{
  max-height: calc(100% - 60px);
}
.el-tabs__content{
  height: calc(100% - 55px);
}
.el-tab-pane{
  height: 100%;
}

.form-buttons{
  display: flex;
  height: 44px;
  border-top: solid 1px #e1e1e1;
  padding-top: 8px;
  align-items: center;
  justify-content: space-around;
  margin: 0px 24px;
}

.left-map,
.right-form {
  transition-duration: 0.5s;
  transition-timing-function: ease-in-out;
  float: left;
}

.left-map {
  width: calc(100% - 581px);
  height: 100%;
}
.right-form {
  width: 70%;
  min-width: 581px;
  max-width: 950px;
}
.no-map .right-form {

  .el-tabs {
    height: calc( 100% - 58px );
    display: flex;
    flex-direction: column;

    .el-tabs__content {
      flex: 1;

      .el-tab-pane {
        height: 100%;
        position: relative;
      }
    }
  }
}

#calendar {
  position: absolute;
  top: 0px;
  left: 23px;
  bottom: 0px;
  right: 23px;
}

.places {
  display: inline;
}

.fc-time-grid-event.fc-v-event.fc-event.fc-start.fc-end{
  //right: 0px;
}
@media (max-width: 1024px) {
  .left-map {
    width: 50%;
  }

  .right-form {
    width: 50%;
    min-width: unset;
  }

  .places {
    display: block;

    label {
      margin-left: 0px;
    }
  }

  .no-map {
    .left-map {
      width: 0%;
    }

    .right-form {
      width: 100%;
    }
  }
}

.content {
  display: flex;
  height: calc(100% - 61px);
}

.small-device {
  align-self: center;
  margin: 0px auto 22px auto;
  font-size: 22px;
}

.fullWidth li {
  padding: 0px;
}

.fullWidth button.el-button {
  width: 100px;
}

.full-height {
  display: flex;

  .el-textarea,
  textarea {
    height: 100%;
  }
}

.flex-none {
  flex: none;
}

.label {
  margin: 0px 15px 0px 19px;
  color: #7d7d7d;
  font-style: 14px;
}

.el-tabs--card>.el-tabs__header {
  .el-tabs__nav {
    border: none;
  }

  .el-tabs__item {
    font-size: 16px;
    font-weight: 600;
    line-height: 2.26;
    letter-spacing: normal;
    text-align: left;
    color: #707070;
    border-color: transparent;
  }

  .el-tabs__item.is-active {
    background-color: #f5f5f5;
    border-color: transparent;
    color: #333333;
    border-bottom: solid 2px #00bff3;
  }
}

.details {
  margin: 24px;

  >* {
    margin-bottom: 14px;
  }
}
.fc-event:not(.fc-helper) {
  height: 22px !important;
  border-radius: 0px;
  margin-left: 0px;
  margin-right: 0px;
  border: 1px solid #00bff3;
}
.info-box{
      margin: 0px;
    padding: 10px 18px;
}
</style>
