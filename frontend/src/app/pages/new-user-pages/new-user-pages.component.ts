import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-new-user-pages',
  templateUrl: './new-user-pages.component.html',
  styleUrls: ['./new-user-pages.component.css']
})
export class NewuserPagesComponent implements OnInit {
  constructor (private route: Router) {}

  private currentuser:string

  private userslist:Array<string> = [];

  ngOnInit(): void {
   
  }

  navigateToHome() {
    this.currentuser=''
    this.route.navigate(['login'])
  }

  setUser(val:string) {
    this.currentuser=val
    this.userslist.push(val)
    //http post
  }
  
  navigateToHomePage(val:string) {
    this.setUser(val)
    this.route.navigate(['home'], {state: {data:val}})
  }
 
}