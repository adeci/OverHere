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

import { BehaviorSubject } from 'rxjs';

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
  private existingPosts:Array<any> = [];
  existingPostsImages:Array<any> = [];
  selectedTags:Array<String> = [];
  imageIDs:Array<any> = [];

  images:Array<any> = [];

  restSelect = '';
  hangSelect = '';
  studySelect = '';
  socialSelect = '';

  responseData$ = new BehaviorSubject<any>(null);

  idk:any[] = [];

  constructor(private markerService: MarkerService, private route: Router, private http: HttpClient, private userservice: UsernameService) { }

  private currentuser = this.userservice.user;
  //variable to keep track of state of post popup
  show = false;
  showkey = false;

  clicked = false;

  

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

  private async initMap(): Promise<void> {
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

    // this.getOHPost().subscribe(
    //   data => {
    //     this.existingPosts = data.data.data;
    //     console.log(this.existingPosts);
    //   },
    //   error => {
    //     console.log("error");
    //   }
    // )

    // this.http.get<any>('http://localhost:8000/ohpost/get/byuserid/' + this.userservice.userid).subscribe(data => {});
    // let ohpostidArr:any[] = [];
    // this.existingPosts = await this.http.get<any>('http://localhost:8000/ohpost/get/byuserid/' + this.userservice.userid).toPromise();
    // let temp:any = Object.values(this.existingPosts)[2].data;
    // for (let i = 0; i < Object.values(this.existingPosts)[2].data.length; i++){
    //   ohpostidArr.push(Object.values(this.existingPosts)[2].data[i].ohpostid);
    // }

    this.images = await this.http.get<any>('http://localhost:8000/images/get/byuserid/' + this.userservice.userid).toPromise();
    console.log(Object.values(this.images)[2].data);
    
    for (let i = 0; i < Object.values(this.images)[2].data.length; i++) {
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
        var newMarker = L.marker([Object.values(this.images)[2].data[i].ycoord, Object.values(this.images)[2].data[i].xcoord], markerOptions);
  
          
        newMarker.bindPopup( 
          "<h1>" + "@" + this.currentuser + 
          "</h1> </div> <img src='" + Object.values(this.images)[2].data[i].encoding + "' width = 200 height = 200 /> <div> <button>Expand</button> </div>"
        );
        newMarker.addTo(this.map);
      }
    }

    //console.log(Object.values(this.existingPosts)[2].data);
    // for (let i = 0; i < Object.values(this.existingPosts)[2].data.length; i++) {
    //     var randIcon = Math.floor(Math.random() * this.pinsList.length);
    //     var randString = Math.floor(Math.random() * this.exampleTags.length);
    //     var randTitle = Math.floor(Math.random() * this.titles.length);
  
    //     var iconProperties: any = {
    //       iconUrl: this.pinsList[randIcon],
    //       iconSize: [38, 45]
    //     }
    //     var customIcon = L.icon(iconProperties);
    //     var markerOptions = {
    //       icon: customIcon,
    //       draggable: false,
    //       title: 'Click to view'
    //     }
    //     var newMarker = L.marker([Object.values(this.existingPosts)[2].data[i].AvgYCoord, Object.values(this.existingPosts)[2].data[i].AvgXCoord], markerOptions);
  
          
    //     newMarker.bindPopup( 
    //       "<h1>" + "@" + this.currentuser + 
    //       "</h1>  <button(click)=\"nagivateToImageView(Object.values(this.existingPosts)[2].data[i].tag, Object.values(this.existingPosts)[2].data[i].caption)\">Expand</button> </div>"
    //     );
    //     newMarker.addTo(this.map);
    //   }

    // let temp:any;
    // let imageArr:any[] = [];
    // for (let i = 0; i < ohpostidArr.length; i++) {
    //   this.images = await this.http.get<any>('http://localhost:8000/images/get/byohpostid/' + ohpostidArr[i]).toPromise();
    //   temp = Object.values(this.images)[2].data;
    //   imageArr.push(temp);
    //   console.log(imageArr);
    // }

    // console.log(imageArr.length);
    // console.log(imageArr[0]);
    // console.log(imageArr[1]);
    // for (let i = 0; i < ohpostidArr.length; i++) {
    //   var randIcon = Math.floor(Math.random() * this.pinsList.length);
    //   var randString = Math.floor(Math.random() * this.exampleTags.length);
    //   var randTitle = Math.floor(Math.random() * this.titles.length);

    //   var iconProperties: any = {
    //     iconUrl: this.pinsList[randIcon],
    //     iconSize: [38, 45]
    //   }
    //   var customIcon = L.icon(iconProperties);
    //   var markerOptions = {
    //     icon: customIcon,
    //     draggable: false,
    //     title: 'Click to view'
    //   }
    //   var newMarker = L.marker([imageArr[i][0].ycoord, imageArr[i][0].xcoord], markerOptions);

        
    //   newMarker.bindPopup( // this.titles[randTitle]
    //     "<h1>" + "@" + this.currentuser + 
    //     "</h1> <div> <p>" +  Object.values(this.existingPosts)[2].data[i].caption + "</p> </div> <div> " + Object.values(this.existingPosts)[2].data[i].tag + " </div> <img src='" + imageArr[i][0].encoding + "' width = 200 height = 200 /> <div> <button>Expand</button> </div>"
    //   );
    //   newMarker.addTo(this.map);
    // }
    //temp =
    

  //}

  setPostsArray() {
    this.getOHPost().subscribe(
      data => {
        this.existingPosts = data.data.data;
      },
      error => {
      }
    )
  }

  ngAfterViewInit(): void {
    //initialize map
    this.initMap();
    //show all of user's posts
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
    ).subscribe((response: any) => {
      console.log(response.data.data.imageid);
      this.imageIDs.push(Object.values(response.data.data.imageid)); 
    });
  }

  // postimgtest(image:String, e:any) {
  //   this.http.post<any>('http://localhost:8000/images/post/', {userid: this.userservice.userid, encoding: image, xcoord: e.latlng.lng, ycoord: e.latlng.lat}).subscribe(response => {
  //     this.userservice.setImgData(response.data.data.imageid);
  //   })
  // }


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

        // this.postimgtest(image, e);
        // const respData = this.userservice.getImgData();
        // console.log(respData);
        // if (respData != "") {
        //   this.imageIDs.push(respData);
        // }

        this.postImage(image, e);
        

        //this.http.post<any>('http://localhost:8000/ohpost/post/withimageids', {userid: this.userservice.userid, tag:tag, caption:caption,imageids: this.imageIDs});

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

  async closeKeyPopup() {
    this.showkey=false;

    let ohpostidArr:any[] = [];
    let tagArr:any[] = [];
    let retArr:any[] = [];

    this.existingPosts = await this.http.get<any>('http://localhost:8000/ohpost/get/byuserid/' + this.userservice.userid).toPromise();
    for (let i = 0; i < Object.values(this.existingPosts)[2].data.length; i++){
      ohpostidArr.push(Object.values(this.existingPosts)[2].data[i].ohpostid);
      tagArr.push(Object.values(this.existingPosts)[2].data[i].tag);
    }

    for (let i = 0; i < tagArr.length; i++){
      if(this.selectedTags.includes(tagArr[i])){
        retArr.push(ohpostidArr[i]);
      }
    }

    //from retArr now i want to simply fetch coordinates. I should have them from the post now

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

  selectRest() {
    this.selectedTags.push("Restaurant");
    this.restSelect = " (Selected)";
  } 
  
  deselectRest() {
    if (this.selectedTags.includes("Restaurants")) {
      var temp = this.selectedTags.indexOf("Restaurants");
      this.selectedTags = this.selectedTags.splice(temp, 1);
      this.restSelect = '';
    }
  }

  selectHang() {
    this.selectedTags.push("Hangout Spot");
    this.hangSelect = " (Selected)"
  } 
  
  deselectHang() {
    if (this.selectedTags.includes("Hangout Spot")) {
      var temp = this.selectedTags.indexOf("Hangout Spot");
      this.selectedTags = this.selectedTags.splice(temp, 1);
      this.hangSelect = '';
    }
  }

  selectStudy() {
    this.selectedTags.push("Study Spot");
    this.studySelect = " (Selected)"
  } 
  
  deselectStudy() {
    if (this.selectedTags.includes("Study Spot")) {
      var temp = this.selectedTags.indexOf("Study Spot");
      this.selectedTags = this.selectedTags.splice(temp, 1);
      this.studySelect = '';
    }
  }

  selectSocial() {
    this.selectedTags.push("Group Meetup/Social Event");
    this.socialSelect = " (Selected";
  } 

  deselectSocial() {
    if (this.selectedTags.includes("Group Meetup/Social Event")) {
      var temp = this.selectedTags.indexOf("Group Meetup/Social Event");
      this.selectedTags = this.selectedTags.splice(temp, 1);
      this.socialSelect = '';
    }
  }

}
