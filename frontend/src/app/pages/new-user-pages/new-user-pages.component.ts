import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { UsernameService } from 'src/app/username.service';

@Component({
  selector: 'app-new-user-pages',
  templateUrl: './new-user-pages.component.html',
  styleUrls: ['./new-user-pages.component.css']
})
export class NewuserPagesComponent implements OnInit {
  constructor (private route: Router, private http: HttpClient, private service: UsernameService) {}

  private currentuser:string = "";

  private userslist:Array<string> = [];

  ngOnInit(): void {
   
  }

  navigateToHome() {
    this.route.navigate(['login'])
  }

  //create a user, post it to the database
  setUser(val:string) {
    this.currentuser=val
    this.userslist.push(val)
    var data = '[{"username": "test1"}]';
    this.http.post('localhost:8000/users/create', data);
    console.log(this.userslist);
    //http post
  }
  
  navigateToHomePage(val:string) {
    this.setUser(val)
    this.service.user = this.currentuser;
    this.route.navigate(['home'], {state: {data:val}})
  }
 
}