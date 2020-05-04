<template>
	<div>
		<el-button v-if="evRanges.length == 0" type="text" size="" v-on:click="addEvRange">{{ not? "Dodaj wykluczone dni":"Dodaj zakres aktywności trasy"}}</el-button>
		<h3 v-if="not && evRanges.length != 0">Wyjątki</h3>
		<el-container v-if="evRanges.length != 0">
			<el-aside width="140" class="left">
				{{not? "Wykluczone dni":"Zakres aktywności trasy"}}
			</el-aside> 
			
			<el-aside width="auto" class="right">
				<div v-for="(item, index) in evRanges" v-bind:data="item" v-bind:key="index">

					<el-date-picker v-if="item.range" v-on:change="change" :picker-options="pickerOptions" v-model="evRanges[index].data" type="daterange" range-separator="do" start-placeholder="Start date" end-placeholder="End date"></el-date-picker>
					<el-date-picker v-if="!item.range" v-on:change="change" :picker-options="pickerOptions" v-model="evRanges[index].data[0]" type="date" start-placeholder="Start date" end-placeholder="End date"></el-date-picker>

					<el-dropdown trigger="click">
						<el-button type="info">
							<i class="el-icon-more"></i>
						</el-button>
						<el-dropdown-menu slot="dropdown" class="fullWidth">
							<el-dropdown-item>
								<el-button :icon="item.range? 'el-icon-time':'el-icon-timer'" size="small" type="text" v-on:click="item.range=!item.range; change();">
									{{item.range? "data":"zakres"}}
								</el-button>
							</el-dropdown-item>
							<el-dropdown-item>
								<el-button icon="el-icon-delete" size="small" type="text" v-on:click="removeEvRange(index)">usuń</el-button>
							</el-dropdown-item>
						</el-dropdown-menu>
					</el-dropdown>
				</div>
				<el-button type="text" size="small" v-on:click="addEvRange">Dodaj kolejne</el-button>
			</el-aside>
		</el-container>
	</div>
</template>

<script>
import moment from 'moment';
export default {
	props: ["data", "open", "not"],
	data() {
		return {
			evRanges: [],
			pickerOptions:{
				disabledDate: this.disabledDate,
				firstDayOfWeek: 1,
			},
		};
	},
	watch:{
		open: function(before, after){
			this.prepareData();
		},
		evRanges: function() { this.change(); }
	},
	methods: {
		disabledDate: function(date){
			return moment(date).isBefore(moment().add(-1, "day"));
		},
		addEvRange() {
			this.evRanges.push({range: true, not: this.not, data: ["",""]});
		},
		removeEvRange(index){
			this.evRanges.splice(index, 1);
			this.change();
		},
		change(){
				//console.log(this);
				let ranges = [];
				for (var i = 0; i < this.evRanges.length; i++) {
					if(this.evRanges[i].data != null)
						ranges.push({
							not: this.evRanges[i].not, 
							data: (this.evRanges[i].range && this.evRanges[i].data.length == 2 && this.evRanges[i].data[1] != "")? 
							[moment(this.evRanges[i].data[0]).format("Y-MM-DD"), moment(this.evRanges[i].data[1]).format("Y-MM-DD ")+"23:59:59"] : 
							[moment(this.evRanges[i].data[0]).format("Y-MM-DD")]});
				}
				this.$listeners['update:data'](ranges);
			},
			prepareData(){
				//console.log(this.data);
				this.evRanges = [];
				for (var i = 0; i < this.data.length; i++) {
					this.evRanges.push({range: this.data[i].data.length == 2, data: this.data[i].data, not: this.data[i].not});
				}
			}
		},
		created: function() { this.prepareData(); }
	}

	</script>

	<style scoped>
	.left{
		margin-top: 5px;
		width: 215px;
		text-align: left;
	}
	.right{
		text-align: left;
		border-left: 1px solid #EBEEF5;
		padding-left: 20px;
	}

	.right>div{
		margin-bottom: 2px;
	}
	.fullWidth li{
		padding: 0px;
	}

	.el-date-editor--daterange.el-input__inner{
		width:400px;
	}

	.fullWidth button.el-button{
		width: 100px;
		text-align: left;
		padding-left: 12px;
	}
	.el-dropdown{
	    top: -1px;
	}

	</style>