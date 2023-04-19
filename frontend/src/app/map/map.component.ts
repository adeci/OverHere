import { Component, AfterViewInit } from '@angular/core';
import * as L from 'leaflet';
import { MarkerService } from '../marker.service';
import { Router } from '@angular/router';
import { IDropdownSettings } from 'ng-multiselect-dropdown/multiselect.model';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { CommonModule, NgIf } from '@angular/common';
import { UsernameService } from 'src/app/username.service';
import { throwError } from 'rxjs';
import { catchError } from 'rxjs';

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
  tempImgID = '';
  existingPosts:Array<any> = [];
  existingPostsImages:Array<any> = [];

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

    this.getOHPost().subscribe(
      data => {
        this.existingPosts = data.data.data.slice(0, data.data.data.length);
        for (let i = 0; i < this.existingPosts.length; i++) {
          this.getImages(this.existingPosts[i].ohpostid).subscribe(
            data => {
              this.existingPostsImages.push(data.data.data);

              var randIcon = Math.floor(Math.random() * this.pinsList.length);
              var randString = Math.floor(Math.random() * this.exampleTags.length);
              var randTitle = Math.floor(Math.random() * this.titles.length);
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
              var newMarker = L.marker([this.existingPostsImages[i][0].ycoord, this.existingPostsImages[i][0].xcoord], markerOptions);
      
              newMarker.bindPopup( // this.titles[randTitle]
                "<h1>" + "@" + this.currentuser + 
                "</h1> <div> <p>" +  this.existingPosts[i].caption + "</p> </div> <div> " + this.existingPosts[i].tag + " </div> <img src='" + this.existingPostsImages[i][0].encoding + "' width = 200 height = 200 /> <div> <button>Expand</button> </div>"
              );
              newMarker.addTo(this.map);
            },
            error => {
              console.log("error");
              return;
            }
          )
        }
        return;
      },
      error => {
        console.log("error");
        
        return;
      }
    )

    
    
    
  }

  ngAfterViewInit(): void {
    //initialize map
    this.initMap();
    //show all of user's posts
  }

  

  getOHPost() {
    return this.http.get<any>('http://localhost:8000/ohpost/get/byuserid/' + this.userservice.userid).pipe(
      catchError((error: HttpErrorResponse) => {
        let errorMessage = 'Unknown error occurred';
        if (error.error instanceof ErrorEvent) {
          // Client-side error
          errorMessage = `Error: ${error.error.message}`;
        } else {
          // Server-side error
          errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
        }
        return throwError(errorMessage);
      })
    );
  }

  getImages(postid:String) {
    return this.http.get<any>('http://localhost:8000/images/get/byohpostid/' + postid).pipe(
      catchError((error: HttpErrorResponse) => {
        let errorMessage = 'Unknown error occurred';
        if (error.error instanceof ErrorEvent) {
          // Client-side error
          errorMessage = `Error: ${error.error.message}`;
        } else {
          // Server-side error
          errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
        }
        return throwError(errorMessage);
      })
    );
  }


  

  postImage(image:String, e:any) {
    return  this.http.post<any>('http://localhost:8000/images/post/', {userid: this.userservice.userid, encoding: image, xcoord: e.latlng.lng, ycoord: e.latlng.lat}).pipe(
      catchError((error: HttpErrorResponse) => {
        let errorMessage = 'Unknown error occurred';
        if (error.error instanceof ErrorEvent) {
          // Client-side error
          errorMessage = `Error: ${error.error.message}`;
        } else {
          // Server-side error
          errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
        }
        return throwError(errorMessage);
      })
    );
  }

  onPlaceClick(image:String, caption:String, tag:String):void {
    if (image === "") {
      image = "./assets/smallspidaman.png";
    }

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

        this.postImage(image, e).subscribe(
          data => {
            var temp:Array<any> = [];
            temp.push(data.data.data.imageid);
            this.http.post<any>('http://localhost:8000/ohpost/post/withimageids', {userid: this.userservice.userid, tag: tag, caption: caption, imageids: temp}).subscribe (data => { });
            return;
          },
          error => {
            console.log("error");
            
            return;
          }
        )

        newMarker.addTo(this.map);
      });
}

  navigateToHomePage() {
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

  