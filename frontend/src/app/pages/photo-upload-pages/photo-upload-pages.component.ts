import { Component, Input } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'photo-upload-pages',
  templateUrl: './photo-upload-pages.component.html',
  styleUrls: ['./photo-upload-pages.component.css']
})
export class PhotoUpPagesComponent {
    constructor (private route: Router) {}
    private photostr: String = "";

  uploadPhoto(val:string) {
    this.photostr = val;
    //http post with image
  }

  navigateToHomePage() {
    this.route.navigate(['home'])
  }


}