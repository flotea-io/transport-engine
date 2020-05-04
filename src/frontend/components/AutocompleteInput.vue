<template>
	<el-container v-loading="loading" >

		<el-popover
		popper-class="autocomplete-popover"
		ref="popover"
		placement="bottom-start"
		trigger="manual"
		:value="showPopover"
		>
		<div class="autocomplete-result" v-for="(result, index) in results" :class="{ 'odd' : isOdd(index)}" @click="select(result)">
			<img :src="'/'+icon">
			<div>
				{{result.label}} 
				<p>
					{{result.region}}
				</p>
			</div>
		</div>
		<div v-if="noResults" class="no-results">No results</div>
		<div></div>
	</el-popover>
	<el-input
	v-popover:popover
	:placeholder="placeholder"
	v-on:focus="focus"
	v-on:blur="blur"
	v-on:clear="clear"
	@keyup.native="keyUp"
	clearable
	v-model="text">
</el-input>
</el-container>
</template>

<script>
import $ from "jquery";
const flotea_api_postal = "https://api.flotea.pl/get-postal";
const flotea_api_postal_flotea = "https://api.flotea.pl/get-postal-flotea";
const flotea_api = "https://geo.flotea.pl/v1/autocomplete";
const ctr_trans = {"DEU":"DE","POL":"PL","FIN":"FI","GRC":"GR","LUX":"LU","NOR":"NO","FRA":"FR","ESP":"ES","MDA":"MD","PRT":"PT","RUS":"RU","SRB":"RS","SCO":"SC","SWE":"SE","SVK":"SK","ITA":"IT","KOS":"KO","BEL":"BE","MKD":"MK","BLR":"BY","ALB":"AL","BGR":"BG","MCO":"MC","AUT":"AT","ROU":"RO","LTU":"LT","LVA":"LV","NIL":"NI","CHE":"CH","SVN":"SI","MLT":"MT","HUN":"HU","CZE":"CZ","BIH":"BA","HRV":"HR","SMR":"SM","DNK":"DK","NLD":"NL","EST":"EE","IRL":"IE","UKR":"UA","WAL":"WA","GBR":"UK","AND":"AD","LIE":"LI"};
export default {
	props: ["value", "placeholder", "icon", "el"],
	components:{
	},
	data: function(){
		return {
			text: "",
			focused: false,
			loading: false,
			noResults: false,
			results: [],
			selected: {},
			disableUpdate: false
		};
	},
	computed: {
		showPopover: function(){
			return this.noResults || this.results.length > 0;
		}
	},
	watch:{
		value: function(newValue){
			this.disableUpdate = true;
			this.text = newValue;
		},
		text: function(newText){
			if(this.disableUpdate){
				this.disableUpdate = false;
				return;
			}
			if(newText.length > 2 && !this.disableUpdate)
				this.fetchData();
		}
	},
	methods: {
		keyUp: function(e){
			this.disableUpdate = false;
			if(e.keyCode == 13 && this.results.length > 0){
				this.select(this.results[0]);
				this.focused = true;
			}
		},
		focus: function(){
			this.disableUpdate = false;
			this.focused = true;
			this.noResults = false;
		},
		blur: function(){
			if(!this.disableUpdate && this.results.length > 0){
				this.select(this.results[0]);
			}
			this.focused = false;
		},
		clear: function(){
			this.$emit('change', {label: ""});
		},
		fetchData: function(){
			this.noResults = false;
			$.ajax({
				url: flotea_api,
				type: "GET",
				data: {
					"text": this.text,
					"layers": "locality,postalcode"
				},
			}).done((respond) => {
				let res = [];
				$( respond.features ).each(( index, item ) => {
					if(typeof item.properties.region != "undefined") {
						res.push({
							label: item.properties.label,
							region: item.properties.region,
							locality: item.properties.locality,
							lat: item.geometry.coordinates[1], 
							lng: item.geometry.coordinates[0],
							country_a: item.properties.country_a
						});                
					}
				});
				this.noResults = res.length == 0;
				this.results = res;
			});
		},
		select: function(item){
			this.selected = item;
			this.results = [];
			this.loading = true;
			$.ajax({
				url: flotea_api_postal_flotea,
				type: "POST",
				data: 'id=' + item.locality +
				'&cnt=' + ctr_trans[item.country_a] +
				'&lat=' + item.lat + '&lng=' + item.lng
			}).done((respond) => {
				this.disableUpdate = true;
				this.loading = false;
				this.selected.zip = respond;
				this.text = this.selected.zip +" "+ this.selected.label;
				this.$emit('change', this.selected);
			})
		},
		isOdd: function(index){
			return index % 2 == 0;
		}
	}
}

</script>

<style lang="scss">
.autocomplete-popover{
	box-shadow: 0 3px 14px #a1a1a1;
	padding: 0px;
	font-family: sans-serif;
	border: 0px;

	.no-results{
		padding: 7px;
		text-align: center;
	}
	.autocomplete-result.odd{
		background: #f1f1f1;
	}
	.autocomplete-result:hover{
		background: #d3e2d3;
		cursor: pointer;
	}
	.autocomplete-result{
		flex-direction: row;
		display: flex;
		font-size: 18px;
		padding: 5px 4px;

		p{
			margin: 0px;
			color: #a1a1a1;
		}
		>img{
			flex: 0;
			align-self: center;
			width: 20px;
			margin-right: 10px;
		}
		>div{
			flex: 1;
		}
	}
}
</style>