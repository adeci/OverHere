import { Component, AfterViewInit } from '@angular/core';
import * as L from 'leaflet';
import { MarkerService } from '../marker.service';
import { Router } from '@angular/router';
import { IDropdownSettings } from 'ng-multiselect-dropdown/multiselect.model';
import { HttpClient } from '@angular/common/http';
import { CommonModule, NgIf } from '@angular/common';
import { UsernameService } from 'src/app/username.service';

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

  // used for testing different dropdown functionalities
  // private dropDownList = [];
  // private selectedItems = [];
  // private dropDownSettings:IDropdownSettings;

  private map: any;
  private marker: any;
  private exampleTags = ['Check this out!', 'Look what I found!', 'Over here!', 'Cute Cat!'];
  private pinsList = ['assets/mapmarkerred.png', 'assets/mapmarkerblue.png', 'assets/mapmarkergreen.png', 'assets/mapmarkerblack.png'];
  private titles = ['test title 1', 'test title 2', 'test title 3'];
  selectedTag = '';
  tempImg = '';

  constructor(private markerService: MarkerService, private route: Router, private http: HttpClient, private userservice: UsernameService) { }

  private currentuser = this.userservice.user;
  //variable to keep track of state of post popup
  show = false;
  showkey = false;

  clicked = false;

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

    // not used with current implentation, saved for other possible implementations of drop down
    // this.dropDownList = [
    //   { item_id : 1 , item_text: 'Restaurant' },
    //   { item_id : 2 , item_text: 'Hangout Spot' },
    //   { item_id : 3 , item_text: 'Study Spot' },
    //   { item_id : 4 , item_text: 'Just for Fun' },
    //   { item_id : 5 , item_text: 'Group Meetup' },
    // ]

    // this.dropDownSettings= {
    //   singleSelection: false,
    //   idField: 'item_id',
    //   textField: 'item_text',
    //   selectAllText: 'Select All',
    //   unSelectAllText: 'UnSelect All',
    //   itemsShowLimit: 3,
    //   allowSearchFilter: false
    // }
  }

  ngAfterViewInit(): void {
    //initialize map
    this.initMap();
    //show all of user's posts
  }

  onPlaceClick(image:String, caption:String, tag:String):void {
    if (image === "") {
      image = "./assets/smallspidaman.png";
    }

    console.log(image);
      (this.map).on('click', (e:any) => {
        var randIcon = Math.floor(Math.random() * this.pinsList.length);
        var randString = Math.floor(Math.random() * this.exampleTags.length);
        var randTitle = Math.floor(Math.random() * this.titles.length);
        //define new map marker object with following properties
        
        var iconProperties: any = {
          iconUrl: this.pinsList[randIcon],
          iconSize: [38, 45]
        }
        var customIcon = L.icon(iconProperties);
        var markerOptions = {
          icon: customIcon,
          draggable: false,
          title: 'Click to view'
        }
        var newMarker = L.marker([e.latlng.lat, e.latlng.lng], markerOptions);

        var c : L.LatLng = newMarker.getLatLng();

        newMarker.bindPopup( // this.titles[randTitle]
          "<h1>" + "@" + this.currentuser + 
          "</h1> <div> <p>" +  caption + "</p> </div> <div> " + tag + " </div> <img src='" + image + "' width = 200 height = 200 /> <div> <button>Expand</button> </div>"
        );
        this.tempImg = '';
        this.http.post<any>('http://localhost:8000/images/post/', {userid: this.userservice.userid, encoding: image, xcoord: e.latlng.lng, ycoord: e.latlng.lat}).subscribe (data => {});
        newMarker.addTo(this.map);
      });
}

  navigateToHomePage() {
    console.log("clicked");
    this.route.navigate(['home']);
  }

  openNewPostPopup() {
    this.show = true;
    this.selectedTag = '';
    this.clicked = false;
  }
  
  openKeyPopup() {
    this.showkey = true;
  }

  closeNewPostPopup() {
    this.show = false;
    this.selectedTag = '';
  }

  closeKeyPopup() {
    this.showkey=false;
  }

  onSelected(value:string): void {
		this.selectedTag = value;
	}

  applyColor(tag:string) {
    //check tag to determine color of pin (for use later)
  }

  createPost(image:String, caption:String, tag:String) {
    this.show = false;
    this.onPlaceClick(image, caption, tag);
  }

  handleUpload(event) {
    const file = event.target.files[0];
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => {
      this.tempImg = <string>reader.result;
    };
  }
}

  