import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-returning-user-pages',
  templateUrl: './returning-user-pages.component.html',
  styleUrls: ['./returning-user-pages.component.css']
})
export class ReturninguserPagesComponent implements OnInit {
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