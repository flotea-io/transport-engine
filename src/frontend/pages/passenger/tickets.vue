<template>
	<div class="container">
		<header-menu activeIndex="2" @metamaskUpdated="(s)=>metamaskState = s" @getObject="(b) => blockchain =b">
		</header-menu>

		<div class="info" v-loading="loading">

			<div v-if="metamaskState.userAddress!='' && !loading && tickets.length == 0" class="text-center">
				<h1>
					Nie masz jeszcze kupionych biletów
				</h1>
				<button @click="goToSearchPage" class="btn-primary mb-4 mt-2">
					Znajdź połączenie
				</button>
				<div class="info-box">
					<b>Wskazówka:</b> aby kupić bilet przejdź do wyszukiwarki połąnczeń wpisz skąd, dokąd chcesz jechać wybierz datę i zapłać tokenem FLT
				</div>
			</div>
		</div>

		<div class="results">
			<div v-for="ticket in tickets" :class="{ended: ended(ticket.Time)}">
				<div class="top">
					Typ pojazdu: {{vehicleTypes[routes[ticket.RouteId].RouteType]}} 
					<span class="company">Firma: <a :href="routes[ticket.RouteId].AgencyUrl">{{routes[ticket.RouteId].AgencyName}}</a></span>
					<span type="text" class="blue-light-button" @click="showMap(ticket.RouteId)">Pokaż mapę</span>

					<span class="time" v-html="getFormatedTime(ticket.Time)"></span>
				</div>

				<h3>
					{{routes[ticket.RouteId].FromLabel}} > {{routes[ticket.RouteId].ToLabel}}
					<span class="price">
						Cena: <span>{{routes[ticket.RouteId].Price}} {{routes[ticket.RouteId].CurrencyType}}</span>
						<br>
						<span>
							{{ticket.Count}} 
						</span>
						{{formatTickets(ticket.Count)}}
					</span>
					<br><small>{{routes[ticket.RouteId].RouteDesc}}</small>
				</h3>

			</div>
		</div>

		<el-dialog width="600px" :visible.sync="mapDialogVisible">
			<div slot="title">
				{{mapDialogTitle}}
				<el-button type="text" icon="el-icon-refresh" @click="centerMap">Wycentruj mapę</el-button>
				<br />
				{{fromLatLng.join(", ")}} &nbsp; - &nbsp; {{toLatLng.join(", ")}}
			</div>
			<client-only>
				<l-map ref="map" @ready="mapReady" :zoom=13 :center="mapPosition">
					<l-tile-layer url="https://maps.heigit.org/openmapsurfer/tiles/roads/webmercator/{z}/{x}/{y}.png"></l-tile-layer>
					<l-marker ref="fromMarker" :lat-lng="fromLatLng">
						<l-icon iconUrl="/icon_from.png" :iconSize="[50, 50]" :iconAnchor="[25, 50]"></l-icon>
					</l-marker>
					<l-marker ref="toMarker" :lat-lng="toLatLng">
						<l-icon iconUrl="/icon_to.png" :iconSize="[50, 50]" :iconAncohr="[25, 50]"></l-icon>
					</l-marker>
				</l-map>
			</client-only>
		</el-dialog>

		<no-metamask-modal text="Aby mieć wgląd we własne bilety, musisz być podłaczony do portfela Medamask. Zaloguj sie lub link." :metamaskDialogVisible.sync="metamaskDialogVisible" />

	</div>
</template>

<script>
import HeaderMenu from "~/components/HeaderMenu.vue";
import NoMetamaskModal from "~/components/NoMetamaskModal.vue";
import moment from 'moment';

export default {
	components: {
		HeaderMenu, NoMetamaskModal
	},
	data: function() {
		return {
			vehicleTypes: ["Tranwaj", "Metro", "Bus", "Prom", "Kolejka linowa", "Gondola", "Kolej linowo-terenowa"],
			blockchain: {},
			routes: {},
			tickets: [],
			mapDialogTitle: "",
			mapDialogVisible: false,
			mapFromLatLng: [],
			mapPosition: [47.413220, -1.219482],
			fromLatLng: [],
			toLatLng: [],
			loading: true,
			metamaskDialogVisible: false,
			metamaskState: {},
		};
	},
	watch: {
      metamaskState: function(n){
        if(!n.enabledMetamask){
          this.metamaskDialogVisible = true;
        } else {
          this.metamaskDialogVisible = false;
          this.loading = false;
          this.loadTickets();
        }
      }
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
	methods: {
		loadTickets: function(){
			let url = process.env.apiUrl+"/v1/search/passager-tickets";
			$.ajax({
				url: url,
				type: "POST",
				data: {wallet: this.metamaskState.userAddress},
			}).done((respond) => {
				for (var i = 0; i < respond.Routes.length; i++) {
					this.routes[respond.Routes[i].RouteId] = respond.Routes[i];
				}
				this.tickets = respond.Tickets;
				this.loading=false;
			});
		},
		getFormatedTime(time){
			let t = moment.unix(time).utc();
			return t.format("D. M. YYYY") +" &nbsp; "+ t.format("H:mm");
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
		showMap: function(index){
			this.fromLatLng = [this.routes[index].FromLat, this.routes[index].FromLng];
			this.toLatLng = [this.routes[index].ToLat, this.routes[index].ToLng];
			this.mapDialogTitle = this.routes[index].RouteDesc;
			this.centerMap();
			this.mapDialogVisible = true;
		},
		centerMap: function(){
			if (typeof this.$refs != "undefined" && typeof this.$refs.map != "undefined")
				this.$refs.map.mapObject.fitBounds([this.fromLatLng, this.toLatLng], { padding: [50, 50] });
		},
		mapReady: function(){
			this.centerMap();
			this.$refs.fromMarker.mapObject.on("click", () => {
				if (!this.dragged)
					this.$refs.map.mapObject.flyTo({ lat: this.fromLatLng[0], lng: this.fromLatLng[1] }, 17, {
						animate: true,
						duration: 1.5
					});
			});
			this.$refs.toMarker.mapObject.on("click", (el) => {
				if (!this.dragged)
					this.$refs.map.mapObject.flyTo({ lat: this.toLatLng[0], lng: this.toLatLng[1] }, 17, {
						animate: true,
						duration: 1.5
					});
			});
		},
		ended: function(time){
			return moment().isAfter(moment.unix(time).utc());
		},
		goToSearchPage: function(){
			location.href = '/';
		}
	}
}
</script>

<style lang="scss">
.container {
	font-family: "Open Sans";
	flex-direction: column;
	align-items: center;
}

.info{
	max-width: 1024px;
	margin: 2em auto;
	position: relative;
	.loading{
		min-height: 220px;
	}
	.el-alert{
		max-width: 34em;
		margin: 0px auto;
	}
}

.results{
	max-width: 1024px;
	margin: 2em auto;
	.ended{
		opacity: 0.4;
	}
	>div{
		margin-bottom: 12px;
		border-radius: 5px;
		box-shadow: 0.5px 0.9px 5px 0 rgba(0, 0, 0, 0.1);
		border: solid 1px #ebebeb;  
		padding: 20px;

		h3{
			margin: 31px 0px 0px;
			font-size: 24px;
			font-weight: 600;
			small{
				color: #707070;
			}
			.price{
				color: #7d7d7d;
				width: 220px;
				float: right;
				text-align: center;
				>span{
					color: #333333;
				}
			}
		}
		.top{
			font-size: 16px;
			color: #707070;
			.blue-light-button{
				background: #e1f5fe;
				border-radius: 5px;
				display: inline-block;
				padding: 1px 12px;
				font-size: 14px;
				font-weight: 600;
				color: #48484a;
			}
			.company{
				margin: 0px 21px;
			}
			a{
				color: #707070;
				font-weight: 600;
			}
			.time{
				color: black;
				font-weight: 400;
				float: right;
				width: 220px;
				text-align: center;
				font-size: 24px;
			}
		}
	}
}

@media (max-width: 768px) {
	.results > div>h3{
		text-align: center;
		.price {
			display: block;
			width: auto;
			float: none;
		}
	}
}


.vue2leaflet-map.leaflet-container{
	padding-bottom: 62%;
}


</style>
