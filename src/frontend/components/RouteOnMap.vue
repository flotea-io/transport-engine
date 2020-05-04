<template>
	<div class="container-block">
		<div class="inputs-group">
			<autocomplete-input placeholder="Skąd" icon="icon_from.png" v-bind:value="from" v-on:change="(data) => change(data, 'from')" /> 
			<autocomplete-input placeholder="Dokąd" icon="icon_to.png" v-bind:value="to" v-on:change="(data) => change(data, 'to')" />
			<section v-if="showMap" class="el-container flex-none">
				<el-button @click="centerMap">Fit map</el-button>
			</section>
			<section class="el-container flex-none">
				<el-checkbox-button v-model="showMap">{{showMap? "Hide":"Show"}} map</el-checkbox-button>
			</section>
		</div>
		<div v-if="showMap" id="map-wrap" style="height: 400px">
			<client-only>
				<l-map ref="map" @ready="mapReady" :zoom=13 :center="mapPosition">
					<l-tile-layer url="https://maps.heigit.org/openmapsurfer/tiles/roads/webmercator/{z}/{x}/{y}.png"></l-tile-layer>

					<l-marker ref="fromMarker" :lat-lng="fromLatLng" @update="(data) => //console.log(data)" :draggable="true">
						<l-icon iconUrl="/icon_from.png" :iconSize="[50, 50]" :iconAnchor="[25, 50]"></l-icon>
					</l-marker>

					<l-marker ref="toMarker" :lat-lng="toLatLng" :draggable="true">
						<l-icon iconUrl="/icon_to.png" :iconSize="[50, 50]" :iconAnchor="[25, 50]"></l-icon>
					</l-marker>
				</l-map>
			</client-only>
		</div>
	</div>
</template>

<script>

import AutocompleteInput from "~/components/AutocompleteInput.vue";
import $ from "jquery";

export default {
	components: {
		AutocompleteInput
	},
	data: function(){
		return {
			from: "",
			to: "",
			mapPosition: [47.413220, -1.219482],
			fromLatLng: [50.05250281819231, 19.98892664909363],
			toLatLng: [52.520856132425436, 13.406206369400024],
			showMap: false,
		};
	},
	watch: {

	},
	created: function(){
		this.dragged = false;
	},
	methods: {
		change: function(val, type){
			this[type+"LatLng"] = [val.lat, val.lng];
			this.centerMap();
		},
		mapReady: function(map){
			this.centerMap();
			this.$refs.fromMarker.mapObject.on("click", ()=> { 
				if(!this.dragged)
					this.$refs.map.mapObject.flyTo({lat: this.fromLatLng[0], lng: this.fromLatLng[1]}, 17);
			});
			this.$refs.toMarker.mapObject.on("click", (el)=> {
				if(!this.dragged)
					this.$refs.map.mapObject.flyTo({lat: this.toLatLng[0], lng: this.toLatLng[1]}, 17);
			});

			this.$refs.fromMarker.mapObject.on('dragend', (data) => {
				this.dragged = true; setInterval(()=> this.dragged = false, 200);
				this.fromLatLng = [data.target._latlng.lat, data.target._latlng.lng];
				this.geoReverse("from", this.fromLatLng);
			});

			this.$refs.toMarker.mapObject.on('dragend', (data) => {
				this.dragged = true; setInterval(()=> this.dragged = false, 200);
				this.toLatLng = [data.target._latlng.lat, data.target._latlng.lng];
				this.geoReverse("to", this.toLatLng);
			});
		},
		geoReverse(type, data){
			let url = "https://geo.flotea.pl/v1/reverse";
			$.ajax({
				url: url,
				type: "GET",
				data: {
					"point.lat": data[0],
					"point.lon": data[1],
				},
			}).done((respond) => {
				if(respond.features.length > 0){
					this[type] = respond.features[0].properties.label;
					//console.log(respond.features[0]);

				}
			});
		},
		centerMap: function() {
			if(typeof this.$refs != "undefined" && typeof this.$refs.map != "undefined")
				this.$refs.map.mapObject.fitBounds([this.fromLatLng, this.toLatLng], { padding: [50, 50] });
		},
	}
}

</script>

<style lang="scss">

</style>
