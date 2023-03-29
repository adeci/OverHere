import { Component } from '@angular/core';
import { Route, Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { UsernameService } from 'src/app/username.service';

@Component({
  selector: 'app-photo-lib-pages',
  templateUrl: './photo-lib-pages.component.html',
  styleUrls: ['./photo-lib-pages.component.css']
})
export class PhotoLibPagesComponent {
  constructor(private route: Router, private http:HttpClient, private service:UsernameService) {}

  private user:String = this.service.user;
  //http get user

  private photos:Array<string> = [];
  //http get photos

  backToHomePage() {
    this.route.navigate(['home'])
  }
}
