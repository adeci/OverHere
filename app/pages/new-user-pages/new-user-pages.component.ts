import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-new-user-pages',
  templateUrl: './new-user-pages.component.html',
  styleUrls: ['./new-user-pages.component.css']
})
export class NewuserPagesComponent implements OnInit {
  constructor (private route: Router) {}

  currentuser:string

  ngOnInit(): void {}

  navigateToHome() {
    this.currentuser=''
    this.route.navigate(['login'])
  }

  
  navigateToHomePage(val:string) {
    //if (this.currentuser.length >=5) {
      this.currentuser=val
      this.route.navigate(['home'], {state: {data:val}})
    //}
  }

  
  
 
}