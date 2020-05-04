<template>
	<div class="center-sm">
		<el-divider content-position="left" v-if="showMore">
			<el-button type="text" icon="el-icon-arrow-left" size="mini" @click="changeMonth(-1)"></el-button>
			<div class="month">
				{{ months[month] }} {{ year }}
			</div>
			<el-button type="text" icon="el-icon-arrow-right" size="mini" @click="changeMonth(1)"></el-button>
		</el-divider>
		<div class="dates">
			<span v-for="d in dates" v-if="showMore || d >= today">
				<span v-bind:class="{selected:day==d, open:eventInDay(d), week:isWeekDay(d)}" 
				@click="selectDateByDay(d)"
				>
				{{d}}
			</span>
		</span>
		<el-button v-if="!showMore" type="text" size="mini" @click="showMore=true">Pokaz kolejne</el-button>
	</div>
	<div class="events">
		<el-date-picker
		class="date-picker"
		size="mini"
		v-model="datePicker"
		type="date"
		:picker-options="pickerOptions"
		value-format="yyyy-MM-dd"
		format="d.M."
		:clearable="false"
		placeholder="Data">
	</el-date-picker>
	<span class="block-sm">
		<label>
			Godzina wyjazdu
		</label>

		<el-select v-model="timeIndex" placeholder="Select" size="mini" :disabled="datePicker==''">
			<el-option
			v-for="(event, index) in eventsInDay"
			:key="index"
			:label="getTimeLabel(event)"
			:value="index">
		</el-option>
	</el-select>
</span>
<span class="block-sm">
	<label>
		Ilość miejsc
	</label>
	<el-input-number v-model="places" size="mini" :min="1" :max="maxTickets" :disabled="datePicker==''"></el-input-number>
</span>

<div class="buy">
	<el-button type="success" :disabled="!metamaskState.enabledMetamask || this.datePicker == ''" size="normal" @click="buyTicket">KUP BILET</el-button>
</div>
</div>
</div>
</template>

<script>
import moment from 'moment';
import timespan from '~/plugins/timespan';
//import { blockChain } from "~/components/Blockchain.vue";
export default {
	props: ["value", "metamaskState", "trip", "blockchain"],
	created: function(){
		this.maxTickets = this.trip.Places;
		this.updateValue(this.value);
		for (var i = 0; i < 12; i++) {
			this.months.push(moment([this.year, i]).format("MMMM"));
		}
	},
	watch: {
		value: function(newValue) { 
			this.updateValue(newValue);
		},
		datePicker: function (newDate) {
			if(newDate == "" || newDate == null) return;
			let d = moment(newDate,"YYYY-M-D");
			this.fillDate(d);
			//this.refreshSchedule();
			if(this.eventInDay(this.day)){
				this.eventsInDay = [];
				for (var i = 0; i < this.events.length; i++) {
					if(this.events[i].open[2]==this.day){
						let open = this.events[i].open.slice(0);
						open[1]--;
						//console.log(this.events[i], open, moment(open).format());
						let close = this.events[i].close.slice(0);
						close[1]--;
						this.eventsInDay.push({
							open: moment.utc(open), 
							close: moment.utc(close)
						});
					}
				}
				this.timeIndex = 0;
				this.changedTime();	
			}
		},
		timeIndex: function(){
			this.changedTime();
		}
	},
	data: function(){
		return {
			showMore: false,
			datePicker: "",
			today: moment().date(),
			day: 1,
			days: 0,
			month: 0,
			year: 2019,
			events: [],
			radio: [],
			eventsInDay: [],
			timeIndex: "",
			pickerOptions:{
				disabledDate: this.disabledDate,
				firstDayOfWeek: 1,
			},
			rules: [],
			months: [],
			places: 1,
			placesText: "",
			buyed: {},
			maxTickets: 100,
			disabledTooltip: true
		};
	},
	computed: {
		dates: function(){
			let r = [];
			for (var i = 1; i <= this.days; i++) {
				r.push(i);
			}
			return r;
		},
	},
	methods: {
		changedTime: function(){
			let buyed = this.buyed[this.eventsInDay[this.timeIndex].open.unix()];
			if(typeof buyed != "undefined"){
				this.maxTickets = this.trip.Places - buyed;
			} else {
				this.maxTickets = this.trip.Places;
			}
		},
		buyTicket: function(){
			if(this.datePicker == ""){
				$.jGrowl("Proszę wybrać datę i godzinę wyjazdu.", { position:"center", header: 'Wskazówka', life: 5000 });
			} else{
				let price = this.trip.Price*this.places*1000;
				if(price > this.metamaskState.flt*1000){
					$.jGrowl("Nie możesz kupić biletu. Masz za mało tokenów FLT.", { position:"center", header: 'Wskazówka', life: 5000 });
					return;
				}
				let count = this.blockchain.numberToHex(this.places); 
				let time = this.blockchain.numberToHex(this.eventsInDay[this.timeIndex].open.unix());
				//let time = this.blockchain.numberToHex(moment("2019-12-17").unix());
				let data = "0x"+'0'.repeat(6-count.length) + count.substr(2) + time.substr(2);
console.log(price, data, this.trip.TripWallet);
				this.blockchain.loadContractAtAddress("Trip", this.trip.TripWallet, ()=>{
					this.blockchain.callContract("Trip"+this.trip.TripWallet, "beforeBuy", (err)=>{
						if(err[0]){
							$.jGrowl(err[1], { position:"center", header: 'Wskazówka', life: 5000 });
						} else {
							this.blockchain.sendContract("FloteaToken", "transfer", { from: this.metamaskState.userAddress }, (a)=>{
								//console.log(a);
								$.jGrowl("Możesz <a href='https://kovan.etherscan.io/tx/"+a+"' target='_blank' class='break-word'>zweryfikować transakcję</a>, lub przejsc na <a href='/passenger/tickets'>liste biletow</a>.", { position:"center", header: 'Sukces', sticky: true });
							}, (err) => {
								//console.log("err",err);
							}, this.trip.TripWallet, price , data);
						}
					}, null, price, data);
				});
			}
		},
		disabledDate: function (date) {
			return moment(date).isBefore(moment().add(-1, "day")) || !timespan.StartInDay(this.rules, moment(date).format("YYYY-M-D").split("-"))
		},
		updateValue: function(newValue){
			let selected = newValue && newValue!=""? moment(newValue) : moment();
			this.fillDate(selected, false);
			//let d = this.day;
			this.rules = timespan.GetRules(JSON.parse(this.trip.Schedule));
			this.refreshSchedule();

			//this.selectDateByDay(d);
		},
		refreshSchedule: function(){
			this.events = [];
			var startDay = timespan.getOnlyDay([this.year, this.month+1, 1]);
			var endDay = timespan.getOnlyDay([this.year, this.month+1, this.days, 23,59,59]);
			var finished = false;
			for (var day = startDay; !finished; day = timespan.getNextDate(day)) {
				this.events = this.events.concat(
					timespan.GetEventsInDay(this.rules, day, 
						[this.year, this.month+1, 1], 
						[this.year, this.month+1, this.days, 23,59,59]
						)
					);
				finished = JSON.stringify(day) == JSON.stringify(endDay);
			}
			this.getTicketsInfo();
		},
		changeMonth: function(month){
			let n = moment([this.year, this.month, 1]).add(month, "month");
			//console.log(n, moment().set({'date': 1, 'hour':0,'minute':0,'second':0}));
			if(n.isBefore(moment().set({'date': 1}).add(-1, "day"))) return;
			this.fillDate(n);
			this.refreshSchedule();
			this.day = 0;
			this.eventsInDay = [];
		},
		selectDateByDay: function(day){
			this.datePicker = moment([this.year, this.month, day]).format('YYYY-MM-DD');
		},
		getTicketsInfo: function(){
			let op = "";
			for (var i = 0; i < this.events.length; i++) {
				let open = this.events[i].open.slice(0);
				open[1]--;
				op += ","+moment.utc(open).unix();
			}
			//console.log(op.substr(1));
			
			let url = process.env.apiUrl+"/v1/search/buyed-tickets-in-times/"+this.trip.RouteId;
			$.ajax({
				url: url,
				type: "POST",
				data: {
					times : op.substr(1)
				}
			}).done((respond) => {
				this.buyed = respond;
			});
		},
		fillDate: function(d, withDay = true){
			if(withDay)
				this.day = d.date();
			this.month = parseInt(d.month());
			this.year = d.year();
			this.days = moment([this.year, this.month]).daysInMonth();
		},
		eventInDay(day){
			for (var i = 0; i < this.events.length; i++) {
				if(this.events[i].open[2]==day)
					return true;
			}
			return false;
		},
		isWeekDay(day){
			let d = moment([this.year, this.month, day]).day();
			return d == 6 || d == 0;
		},
		getTimeLabel(event){

			let buyed = this.buyed[event.open.unix()];
			if(typeof buyed != "undefined"){
				return event.open.format('H:mm') + " - " + (this.trip.Places-buyed) + " " + this.formatTickets(this.trip.Places-buyed);
			} else {
				return event.open.format('H:mm') + " - " + this.trip.Places + " " + this.formatTickets(this.trip.Places);
			}

		},
		formatTickets(tickets){
			if(tickets == 1){
				return "bilet";
			}
			if(tickets < 5){
				return "bilety";
			}
			return "biletów";
		},
	}
}

</script>

<style lang="scss">
.month{
	width: 126px;
	color: #48484a;
	font-size: 14px;
	display: inline-block;
	text-align: center;
	font-weight: 600;
}
.el-divider__text .el-button{
	color: #333333;
	padding: 0px 6px;
	margin: 0px -8px;
}
.el-divider{
	width: calc(100% - 220px);
}
.events{
	margin-top:1em;
	color: #7d7d7d;
	font-size: 14px;
	.el-select{
		width: 150px;
	} 
	.date-picker{
		width: 100px;
	}
	.el-select{
		margin: 0px 37px 0px 6px;
	}
	.date-picker{
		margin-right: 42px;
	}
	.el-input-number{
		margin: 6px;
	}
	.buy{
		width: 220px;
		float: right;
		text-align: center;
		>.el-button{
			border-radius: 5px;
			background-color: #00bff3;
			font-size: 14px;
			font-weight: bold;
			color: #fff;
			border-color: #00bff3;
			&.is-disabled{
				border-color: #d6d6d6;
				background-color: #d6d6d6;
				color: #9c9c9c;
			}
		}
	}
}

.dates{
	margin-bottom: 8px;
	display: flex;
	width: calc(100% - 220px);
	>span{
		flex: none;
		>span{
			background-color: #ffffff;
			color: #464646;
			border-radius: 100%;
			height: 22px;
			width: 22px;
			display: inline-block;
			text-align: center;
			font-size: 14px;
			line-height: 22px;
			font-weight: 600;

			&.week{
				background-color: #ebebeb;
			}
			&.open{
				cursor: pointer;
				background-color: #e1f5fe;
				&:hover, &.selected{
					color: white;
					background-color: #00bff3;
				}
			}
		}
	}
}
.jGrowl-message a{
	color: white;
}
@media (max-width: 978px) {
	.el-divider{
		display: none;
	}
	.dates{
		display: none;
	}
	.events .buy {
		width: auto;
		display: inline-block;
		float: none;
	}
}
@media (max-width: 768px){
	.events{
		text-align: center;
		.date-picker{
			margin-right: 0px;
			margin-bottom: 6px;
		}
		.buy{
			display: block;
		}
	} 
	.block-sm{
		display: block;
		margin-bottom: 6px;
	}
	.center-sm{
		label{
			width: 118px;
			display: inline-block;
			text-align: right;
			padding-right: 6px;
		}
		.el-select{
			margin: 0px;
			width: 130px;
		}
		.el-input-number{
			margin: 0px;
			width: 130px;
		}
	}
}
</style>