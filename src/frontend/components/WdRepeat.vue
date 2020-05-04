<template>
	<el-container>
		<el-aside width="140" class="left">
			<el-radio-group v-model="allWeek" size="small">
				<el-radio-button label="0">Cały tydzień</el-radio-button>
				<el-radio-button label="1">Dni tygodnia</el-radio-button>
			</el-radio-group>
		</el-aside>
		<el-aside class="right" width="calc(100% - 215px)">
			<el-select v-if="allWeek=='1'" size="small" v-model="checkedWd" @change="change" multiple placeholder="Wybierz dni">
				<el-option
				v-for="(day, index) in weekDays"
				:key="index"
				:label="day"
				:value="day">
			</el-option>
		</el-select>

		<!--<el-checkbox v-for="(day, index) in weekDays" v-bind:data="day" @change="change" v-bind:key="index" v-model="checkedWd[index]">{{day}}</el-checkbox>-->
	</el-aside>
</el-container>
</template>

<script>
import moment from 'moment';
import "moment/locale/pl";
export default {
	props: ["data","open"],
	data() {
		return {
			checkedWd: [],
			weekDays: [],
			allWeek: "0"
		};
	},
	watch:{
		open: function(before, after){
			this.prepareData();
		},
		allWeek: function(selected){
			if(selected == "0")
				this.checkedWd = [];
		}
	},
	methods: {
		change(){
			let indexes = [];
			for (var i = 0; i < this.checkedWd.length; i++) {
				indexes.push(this.weekDays.indexOf(this.checkedWd[i])+1);
			}
			this.$listeners['update:data'](indexes);
		},
		prepareData(){
			this.checkedWd = [];
			for (var i = 1; i <= 7; i++) {
				if(this.data.indexOf(i) != -1 ){
					this.checkedWd.push(this.weekDays[i-1]);
				}
			}
		}
	},
	created: function() { 
		moment().locale("pl");
		for (var i = 1; i <= 7; i++) {
			this.weekDays.push(moment("2019-07-0"+i).format('dddd', "pl"));
		}
		this.prepareData(); 
	}
}
</script>

<style scoped>
.left{
	width: 215px;
}
.right{
	text-align: left;
	padding-left: 21px;
	border-left: 1px solid #EBEEF5;
}
.el-checkbox{
	margin-right: 14px;
}
.right > .el-select{
	display: block;
	overflow: hidden;
}
</style>