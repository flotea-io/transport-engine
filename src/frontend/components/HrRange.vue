
<template>
	<el-container class="HrRange">
		<el-aside width="140" class="left">
			Godzina wyjazdu
		</el-aside> 
		<el-aside width="auto" class="right">
			<div v-for="(item, index) in hrRanges" v-bind:data="item" v-bind:key="index">
				<el-time-picker
				v-on:change="changeFrom"
				format="H:mm"
				v-model="hrRanges[index].open"
				:clearable="false"
				placeholder="Start">
			</el-time-picker>
			do
			<el-select v-model="hrRanges[index].days" v-on:change="change">
				<el-option
				v-for="(item, i) in days"
				:key="i"
				:label="item"
				:value="i">
			</el-option>
		</el-select>
		<el-time-picker
		format="H:mm"
		:picker-options="{ selectableRange: minTime(index) }"
		v-on:change="change"
		v-model="hrRanges[index].close"
		:clearable="false"
		placeholder="End">
	</el-time-picker>
					<el-button icon="el-icon-delete"  type="info" v-on:click="removeHrRange(index)"></el-button>

</div>
<el-button type="text" size="small" v-on:click="addHrRange">Dodaj kolejne</el-button>
</el-aside>
</el-container>
</template>

<script>
	import moment from 'moment';
	export default {
		props: ["data", "open"],
		data: function(){
			return {
				days: ["w tym dniu", "następny dzień", "+ 2 dni", "+ 3 dni", "+ 4 dni", "+ 5 dni", "+ 6 dni"],
				hrRanges: []
			};
		},
		watch:{
			open: function(before, after){
				this.prepareData();
			},
			hrRanges: function() { this.change(); }
		},
		methods: {
			minTime(index){
				if(this.hrRanges[index].days>0)
					return "00:00:00 - 23:59:59";
				let d = moment(this.hrRanges[index].open).add(1,'minutes');
				return d.format("HH:mm:ss")+" - 23:59:59";
			},
			addHrRange() {
				this.hrRanges.push({days: 0, open: "2019-01-01 08:00:00", close: "2019-01-01 09:00:00"});
			},
			removeHrRange(index){
				if(this.hrRanges.length > 1){
					this.hrRanges.splice(index, 1);
				}
				else{
					this.$alert('Must be minimal one', 'Warning', {
          				confirmButtonText: 'OK',
        			});
				}
			},
			changeFrom(){
				for (var i = 0; i < this.hrRanges.length; i++) {
					let open = moment(this.hrRanges[i].open);
					if(moment(this.hrRanges[i].close).isBefore(open) && this.hrRanges[i].days == 0){
						this.hrRanges[i].close = open.clone().add(1,'minutes').format("Y-MM-DD HH:mm:ss");
					}
				}
				this.change();
			},
			change(){
				let ranges = [];
				for (var i = 0; i < this.hrRanges.length; i++) {
					let open = moment(this.hrRanges[i].open);
					let close = moment(this.hrRanges[i].close);

					let hr = close.diff(open, "hours");
					hr -= 24 * Math.floor(hr/24);
					close = close.format("HH.mm.ss").split(".").map(numStr => parseInt(numStr));
					close[0] = hr + 24*this.hrRanges[i].days ;
					ranges.push([open.format("HH.mm.ss").split(".").map(numStr => parseInt(numStr)), close]);
				}
				this.$listeners['update:data'](ranges);
			},
			prepareData(){
				this.hrRanges = [];
				let time = ["h", "m", "s"];
				for (var i = 0; i < this.data.length; i++) {
					let open = moment([2019,0,1].concat(this.data[i][0]));
					let close = open.clone();
					for (var j = 0; j < this.data[i][1].length; j++) {
						close.add(this.data[i][1][j], time[j]);
					}
					this.hrRanges.push({days: close.diff(open, 'days'), open: open.format("Y-MM-DD HH:mm:ss"), close: close.format("Y-MM-DD HH:mm:ss")});
				}	
			}
		},
		created: function() { this.prepareData(); }
	}

</script>
<style>
.HrRange .el-date-editor input {
		padding-right: 0px!important;
}
</style>
<style scoped>
.el-select{
	width: 116px;
}
.el-date-editor.el-input{
	width: 82px;

}

.el-aside.right{
	overflow: hidden;
}
.right>div{
		margin-bottom: 3px;
	}
.left{
	width: 215px;
	margin-top: 5px;
	text-align: left;
}
.right{
	text-align: left;
	border-left: 1px solid #EBEEF5;
	padding-left: 20px;
}
.fullWidth li{
	padding: 0px;
}

.fullWidth button.el-button{
	width: 100px;
	text-align: left;
	padding-left: 12px;
}
</style>