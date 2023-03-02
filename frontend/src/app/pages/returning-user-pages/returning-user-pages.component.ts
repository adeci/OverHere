import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RouterModule } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { UsernameService } from 'src/app/username.service';

@Component({
  selector: 'app-returning-user-pages',
  templateUrl: './returning-user-pages.component.html',
  styleUrls: ['./returning-user-pages.component.css']
})
export class ReturninguserPagesComponent implements OnInit {
  constructor (private route: Router, private http: HttpClient, private service: UsernameService) {}

  currentuser:string = ""

  ngOnInit(): void {}

  navigateToHome() {
    this.currentuser=''
    this.route.navigate(['login'])
  }

  
  navigateToHomePage(val:string) {
    this.currentuser=val
    this.service.user = this.currentuser;
    this.route.navigate(['home'], {state: {data:val}})
  }

}