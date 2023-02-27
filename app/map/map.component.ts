import { Component, AfterViewInit } from '@angular/core';
import * as L from 'leaflet';
import { MarkerService } from '../marker.service';
//leaflet and angular link code based on tutorial posted by digital ocean
// https://www.digitalocean.com/community/tutorials/angular-angular-and-leaflet

const iconRetinaUrl = 'assets/marker-icon-2x.png';
const iconUrl = 'assets/marker-icon.png';
const shadowUrl = 'assets/marker-shadow.png';
const iconDefault = L.icon({
  iconRetinaUrl,
  iconUrl,
  shadowUrl,
  iconSize: [25, 41],
  iconAnchor: [12, 41],
  popupAnchor: [1, -34],
  tooltipAnchor: [16, -28],
  shadowSize: [41, 41]
});
L.Marker.prototype.options.icon = iconDefault;

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.css']
})
export class MapComponent implements AfterViewInit {
  private map: any;
  private marker: any;
  private exampleTags = ['Check this out!', 'Look what I found!', 'Over here!', 'Cute Cat!'];
  private pinsList = ['assets/mapmarkerred.png', 'assets/mapmarkerblue.png', 'assets/mapmarkergreen.png', 'assets/mapmarkerblack.png'];

  constructor(private markerService: MarkerService) { }

  private initMap(): void {
    //initialize center point (Ben Hill Griffin Stadium, Gainesville FL used as test)
    this.map = L.map('map', {
      center: [ 29.649934, -82.348655 ],
      zoom: 15
    });

    //intanstiatie map tile layer- uses openstreet map image/API
    const tiles = L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      maxZoom: 18,
      minZoom: 3,
      attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    });

    tiles.addTo(this.map);
  }

  ngAfterViewInit(): void {
    //initialize map
    this.initMap();
    //place a marker
    this.map.on('click', (e: any) => {
      var randIcon = Math.floor(Math.random() * this.pinsList.length);
      var randString = Math.floor(Math.random() * this.exampleTags.length);
      //define new map marker object with following properties
      var iconProperties: any = {
        iconUrl: this.pinsList[randIcon],
        iconSize: [38, 45]
      }
      var customIcon = L.icon(iconProperties);
      var markerOptions = {
        icon: customIcon,
        draggable: false
      }
      var newMarker = L.marker([e.latlng.lat, e.latlng.lng], markerOptions);

      //this.exampleTags[randString]
      var c : L.LatLng = newMarker.getLatLng();

      newMarker.bindPopup("Lat: " + c.lat + ", Lng: " + c.lng, {
        closeButton: true
      })
      newMarker.addTo(this.map);
    });
    

  }
}
