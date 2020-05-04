import moment from 'moment';

var scheduler = {
	data: true,
	removeZeroInArray(_arr, month){
		let arr = _arr.slice(0);
		if(month && arr.length > 1){
			arr[1] +=1;
		}
		for (var i = arr.length - 1; i >= 0; i--) {
			if(arr[i] != 0)
				return arr.slice(0,i+1);
		}
	},
	arrayDateToString(_arr, addMonth = false){
		let arr = _arr.slice(0);
		if( arr.length > 1)
			arr[1]--;
		return moment(arr).format("Y-MM-DD HH:mm:ss");
	},
	arrayHourToString(arr, full){
		let z = "2020-01-01 "+
		(typeof arr[0] != "undefined"? (arr[0]<10? "0"+arr[0] : arr[0]) : "00")+":"+
		(typeof arr[1] != "undefined"? (arr[1]<10? "0"+arr[1] : arr[1]) : "00")+":"+
		(typeof arr[2] != "undefined"? (arr[2]<10? "0"+arr[2] : arr[2]) : "00");
		//console.log(z, moment(z), moment(z).format("HH:mm:ss"));
		if(full)
			return moment(z).format("Y-MM-DD HH:mm:ss");
		else
			return moment(z).format("HH:mm:ss");
	},

	events: [],
	reset(){
		this.events = [];
	},
	resolve(id, testedDay, o){
		this.id = id;
		//console.log(id);
		this.testedDay = testedDay;
		this.weekDay = testedDay.weekday();
		//console.log(testedDay, o);
		this.resolve2(o, false);
	},
	
	resolve2(o, depend) {
		let returnObject = {ok: true};
		if(typeof o.or != "undefined"){ // Array of non depended
			let results = [];
			for (var i = 0; i < o.or.length; i++) {
				if(depend){
					let z = this.resolve2(o.or[i], false || depend);
					results.push(z);
				}
				else 
					this.resolve2(o.or[i], false || depend);
			}
			if(depend) {
				let onlyB = this.onlyBoolean(results);
				if(onlyB.ok) 
					return [onlyB];

				return results;
			}
			return;
		}
		if(typeof o.and != "undefined"){ // Array of depended
			let results = []; 
			for (var i = 0; i < o.and.length; i++) {
				let z = this.resolve2(o.and[i], true || depend);
				if(Array.isArray(z)){
					for (var j = 0; j < z.length; j++) {
						results.push(z[j]);        
					}
				} else results.push(z);
			}
			if(depend) return returnObject;

			this.testResolve(results);
			return;
		}
		if(typeof o.not != "undefined"){
			let z = this.resolve2(o.not, true || depend);
			z.ok = !z.ok;
			return [z];
		}
		if(typeof o.wd != "undefined"){
			returnObject.ok &= o.wd.indexOf(this.weekDay) != -1;
		}
		if(typeof o.bdate != "undefined"){
			if(o.bdate.length == 0 || o.bdate.length > 2) console.error("Wrong bdate format");
			returnObject.ok &= (o.bdate[0] <= this.testedDay.month()+1);
			if(o.bdate.length == 2)
				returnObject.ok &= (o.bdate[1] <= this.testedDay.date());
		}
		if(typeof o.edate != "undefined"){
			if(o.edate.length == 0 || o.edate.length > 2) console.error("Wrong edate format");
			returnObject.ok &= (o.edate[0] >= this.testedDay.month()+1);
			if(o.edate.length == 2)
				returnObject.ok &= (o.edate[1] >= this.testedDay.date());
		}
		if(typeof o.bev != "undefined"){
			if(o.bev.length == 0 || o.bev.length > 6) console.error("Wrong bev format");
			let bevMoment = this.getMomentFromEv(o.bev, true, false);
			returnObject.ok = returnObject.ok && bevMoment.unix() <= this.testedDay.unix();
			if(o.bev.length > 3 && bevMoment.isSame(this.testedDay))
				returnObject.bev = this.getMomentFromEv(o.bev, false, true);

		}
		if(typeof o.eev != "undefined"){
			if(o.eev.length == 0 || o.eev.length > 6) console.error("Wrong eev format");
			let eevMoment = this.getMomentFromEv(o.eev, true, false);
			//console.log(eevMoment, eevMoment.unix(), this.testedDay, this.testedDay.unix());
			returnObject.ok = returnObject.ok && eevMoment.unix() >= this.testedDay.unix();
			if(o.eev.length > 3 && eevMoment.isSame(this.testedDay))
				returnObject.eev = this.getMomentFromEv(o.eev, false, true);
		}

		if(typeof o.bhr != "undefined"){
			returnObject.bhr = o.bhr;
		}
		if(typeof o.ehr != "undefined"){
			returnObject.ehr = o.ehr;
		}


		if(depend) return returnObject;
		else {
			this.testResolve(returnObject);
		}
	},

	getHourDepend(arr) {
		let hourDepend = [];
		if(Array.isArray(arr)){
			for (var i = 0; i < arr.length; i++) {
				let obj = {};
				if(typeof arr[i].bev != "undefined")
					obj.bev = arr[i].bev;
				if(typeof arr[i].eev != "undefined")
					obj.eev = arr[i].eev;
				if(Object.keys(obj).length > 0)
					hourDepend.push(obj);
			}
		}
		return hourDepend;
	},

	testResolve(arr) {
		if(Array.isArray(arr)){
			let hourDepend = this.getHourDepend(arr);
			for (var i = 0; i < arr.length; i++) {
				if(typeof arr[i].ok == "boolean" && arr[i].ok == false)
					return false;
				if(Array.isArray(arr[i]) && typeof arr[i].find((el)=> { return typeof el.ok != "undefined" && el.ok == false; }) != "undefined")
					return false;
				if(hourDepend.length > 0)
					arr[i].hourDepend = hourDepend;
			}
			this.createEvents(arr);
			return true;
		}else{
			let keys = Object.keys(arr);
			for (var i = 0; i < keys.length; i++) {
				if(keys[i] == "ok" && arr[keys[i]] == false)
					return false;
			}

			let mini = true;
			if(typeof arr.bev != "undefined"){
				mini &= arr.bev <= this.arrayToSeconds(arr.bhr);
			}
			if(typeof arr.eev != "undefined"){
				mini &= arr.eev >= this.arrayToSeconds(arr.ehr);
			}
			if(!mini) return false;
			this.createEvents(arr);
			return true;
		}
	},

	createEvents(arr){
		if(Array.isArray(arr)){
			for (var i = 0; i < arr.length; i++) {
				if(typeof arr[i].bhr == "undefined" && typeof arr[i].ehr == "undefined") continue;
				if(this.testDependHours(arr[i])) continue;
				this.events.push({
					title: "Plan wyjazdu #" + (this.id+1),
					myId: this.id,
					start: typeof arr[i].bhr == "undefined"? this.testedDay : this.getMoment(moment(this.testedDay), arr[i].bhr),
					//end: typeof arr[i].ehr == "undefined"? moment(this.testedDay).add(1,"h") : this.getMoment(moment(this.testedDay), arr[i].ehr),
				});
			}
		}else{
			this.events.push({
				title: "Plan wyjazdu #" + (this.id+1),
				myId: this.id,
				start: typeof arr.bhr == "undefined"? this.testedDay : this.getMoment(moment(this.testedDay), arr.bhr),
				//end: typeof arr.ehr == "undefined"? moment(this.testedDay).add(1,"h") : this.getMoment(moment(this.testedDay), arr.ehr),
			});
		}
	},

	testDependHours(arr){
		if(typeof arr.hourDepend == "undefined") return false;
		for (var i = 0; i < arr.hourDepend.length; i++) {
			let mini = true;
			if(typeof arr.hourDepend[i].bev != "undefined"){
				mini &= arr.hourDepend[i].bev <= this.arrayToSeconds(arr.bhr);
			}
			if(typeof arr.hourDepend[i].eev != "undefined"){
				mini &= arr.hourDepend[i].eev >= this.arrayToSeconds(arr.ehr);
			}
			if(mini) return false;
		}
		return true;
	},

	getMomentFromEv(arr, onlyDay, seconds){
		let z = arr[0]+"-"+(typeof arr[1] != "undefined"? (arr[1]<10? "0"+arr[1] : arr[1]) : "01")+
		"-"+ (typeof arr[2] != "undefined"? (arr[2]<10? "0"+arr[2] : arr[2]) : "01");
		if(onlyDay) return moment.utc(z);
		z += " "+
		(typeof arr[3] != "undefined"? (arr[3]<10? "0"+arr[3] : arr[3]) : "00")+":"+
		(typeof arr[4] != "undefined"? (arr[4]<10? "0"+arr[4] : arr[4]) : "00")+":"+
		(typeof arr[5] != "undefined"? (arr[5]<10? "0"+arr[5] : arr[5]) : "00");
    	if(seconds) return this.momentTimeToSeconds(moment.utc(z));
    		return moment.utc(z);
	},

	momentTimeToSeconds(m){
		return m.hours() * 3600 + m.minutes() * 60 + m.seconds();
	},

	arrayToSeconds(arr){
		return (typeof arr[0] == "undefined"? 0 : arr[0] * 3600) +
		(typeof arr[1] == "undefined"? 0 : arr[1] * 60) +
		(typeof arr[2] == "undefined"? 0 : arr[2]);
	},

	getMoment(m, obj){
		let units = ["h", "m", "s"]
		for (var i = 0; i < units.length; i++) {
			if(obj[i] != "undefined")
				m.add( obj[i], units[i]);
		}
		return m;
	},

	onlyBoolean(arr){
		let obj = { "ok" : this.findTrue(arr) };
		for (var i = 0; i < arr.length; i++) {
			if(typeof arr[i].bhr != "undefined" || typeof arr[i].ehr != "undefined") return { "ok" : false }; 
			if(typeof arr[i].bev != "undefined") obj.bev = arr[i].bev;
			if(typeof arr[i].eev != "undefined") obj.eev = arr[i].eev;
		}
		return obj;
	},

	findTrue(arr){
		for (var i = 0; i < arr.length; i++) {
			if(arr[i].ok) return true; 
		}
		return false;
	}


};	

export default scheduler;