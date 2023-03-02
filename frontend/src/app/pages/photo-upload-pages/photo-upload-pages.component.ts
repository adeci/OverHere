import { Component, Input } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'photo-upload-pages',
  templateUrl: './photo-upload-pages.component.html',
  styleUrls: ['./photo-upload-pages.component.css']
})
export class PhotoUpPagesComponent {
    constructor (private route: Router, private http:HttpClient) {}
    private photostr: String = "";
    private user='';
    //http get user

  uploadPhoto(val:string) {
    this.photostr = val;
    var data = '{["imageid":"sentimg", "userid":"sentimg2", "ohpostid":"sentimg3", "encoding":"sentimg4"]}';
    this.http.post('localhost:8000/images/create', data);
    //http post with image (test for demo)
  }

  navigateToHomePage() {
    this.route.navigate(['home'])
  }


}