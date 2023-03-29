import { Component, Input } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { UsernameService } from 'src/app/username.service';

@Component({
  selector: 'photo-upload-pages',
  templateUrl: './photo-upload-pages.component.html',
  styleUrls: ['./photo-upload-pages.component.css']
})
export class PhotoUpPagesComponent {
    constructor (private route: Router, private http:HttpClient, private service:UsernameService) {}
    private photostr: String = "";
    private user:String=this.service.user;
    //http get user

  uploadPhoto(val:String) {
    this.photostr = val;
    console.log(this.service.user);
    this.http.post<any>('http://localhost:8000/images/create', {imageid: "hello", userid: <string>(this.service.user), ohpostid: "number", encoding: "code"}).subscribe (data => {});
  }

  navigateToHomePage() {
    this.route.navigate(['home'])
  }


}