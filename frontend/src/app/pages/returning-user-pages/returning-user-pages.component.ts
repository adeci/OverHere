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

  validuser=true;

  currentuser:string = "";
  getuser:string = "";

  ngOnInit(): void {}

  navigateToHome() {
    this.currentuser=''
    this.route.navigate(['login'])
  }

  validateUser() {
    //http get
    // if valid, navigate to homepage
  }
  
  navigateToHomePage(val:string) {
    this.currentuser=val
    this.service.user = this.currentuser;
    //this.http.get<string>('http://localhost:8000/users/get/123456').subscribe(data => {console.log(data)});
    
    this.route.navigate(['home'], {state: {data:val}})
  }

}