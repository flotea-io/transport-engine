import moment from 'moment';

export default class Event {
	constructor(id, event){
		console.log(event);
		this.id = id;
		if(event.allDay){
			this.calendarStart = event.start.set({'hours': 8,'minutes': 0, 'seconds': 0});
			this.calendarEnd = this.calendarStart.clone().add(1, "h");//.add(1, "d");
		} else {
			this.calendarStart = event.start;//.add(50,"s");
			this.calendarEnd = event.end;//.add(1,"d");
        }
        this.checkedWd = [];

        let minDiff = this.calendarEnd.minute() - this.calendarStart.minute();
        let secDiff = this.calendarEnd.second() - this.calendarStart.second();

        if (secDiff < 0) {
            secDiff = 60 + secDiff;
            minDiff--;
        }
        if (minDiff < 0) {
            minDiff = 60 + minDiff;
        }

        this.bhr = [this.removeZeroInArray(this.calendarStart.toArray().slice(3,6))];
        this.hr =  [[moment.duration(this.calendarEnd.diff(this.calendarStart)).hours(),minDiff,secDiff]];
        this.bev = [this.removeZeroInArray(this.calendarStart.toArray().slice(0,3), true)];
        this.eev = [this.removeZeroInArray(this.calendarEnd.toArray().slice(0,3), true)];
        //console.log(this.calendarEnd.toArray().slice(0,3).concat([23,59,59]));
        this.object = {
       		"bev": this.removeZeroInArray(this.calendarStart.toArray().slice(0,3), true),
       		"eev": this.removeZeroInArray(this.calendarEnd.toArray().slice(0,3).concat(23,59,59), true),
       		"bhr": this.removeZeroInArray(this.bhr[0]),
       		"hr": this.removeZeroInArray(this.hr[0])
        };
       
        this.modal = {
        	id: this.id,
        	dialogVisible: false,
        	checkedWd: this.checkedWd,
        	evRange:[{not: false, data:[this.arrayDateToString(this.bev[0]), this.arrayDateToString(this.eev[0])]}],
        	notEvRange:[],
        	hrRange:[[this.object.bhr, this.object.hr]],
        	mdRange:[]
        };
	}


	updateFromModal(){
		this.checkedWd = this.modal.checkedWd;
		let object = {"and":[]};

		if(this.modal.checkedWd.length > 0){
			object.and.push({ "wd": this.modal.checkedWd });
		}
		if(this.modal.evRange.length > 0){
			let r = {"or":[]};
			for (var i = 0; i < this.modal.evRange.length; i++) {
				let re = {};
				re.bev = this.removeZeroInArray(moment(this.modal.evRange[i].data[0]).toArray(), true);
				if(this.modal.evRange[i].data.length > 1){
					re.eev = this.removeZeroInArray(moment(this.modal.evRange[i].data[1]).toArray(), true)
				}
				r.or.push(re);
			}
			object.and.push(r);
		}

		if(this.modal.notEvRange.length > 0){
			let n = {"not":[]};
			for (var i = 0; i < this.modal.notEvRange.length; i++) {
				let re = {};
				re.bev = this.removeZeroInArray(moment(this.modal.notEvRange[i].data[0]).toArray(), true);
				if(this.modal.notEvRange[i].data.length > 1){
					re.eev = this.removeZeroInArray(moment(this.modal.notEvRange[i].data[1]).toArray(), true)
				}
				n.not.push(re);
			}
			object.and.push(n);
		}

		if(this.modal.hrRange.length > 0){
			let r = {"or":[]};
			for (var i = 0; i < this.modal.hrRange.length; i++) {
				let re = {
					bhr: this.removeZeroInArray(this.modal.hrRange[i][0]),
					hr: this.removeZeroInArray(this.modal.hrRange[i][1])
				};
				r.or.push(re);
			}
			object.and.push(r);
		}

		/*if(this.modal.mdRange.length > 0){
			let r = {"or":[]};
			for (var i = 0; i < this.modal.mdRange.length; i++) {
				let re = {};
				re.bdate = this.modal.mdRange[i].data[0];
				if(this.modal.mdRange[i].data.length > 1){
					re.edate = this.modal.mdRange[i].data[1];
				}
				if(this.modal.mdRange[i].not)
					r.or.push({"not": re });
				else
					r.or.push(re);
			}
			object.and.push(r);
		}*/
		//console.log(object);
		this.object = object;
		//console.log(JSON.stringify(this.object, null, 3));

		//console.log(JSON.stringify(this, null, 3));
	}

	removeZeroInArray(_arr, month){
		let arr = _arr.slice(0);
		if(month && arr.length > 1){
			arr[1] +=1;
		}
		for (var i = arr.length - 1; i >= 0; i--) {
			if(arr[i] != 0)
				return arr.slice(0,i+1);
		}
	}

	arrayDateToString(_arr, addMonth = false){
		let arr = _arr.slice(0);
		if( arr.length > 1)
			arr[1]--;
		return moment(arr).format("Y-MM-DD HH:mm:ss");
	}
}