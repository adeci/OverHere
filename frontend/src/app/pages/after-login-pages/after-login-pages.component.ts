import { Component, Input } from '@angular/core';
import { Router } from '@angular/router';

import { UsernameService } from 'src/app/username.service';

@Component({
  selector: 'app-after-login-pages',
  templateUrl: './after-login-pages.component.html',
  styleUrls: ['./after-login-pages.component.css']
})
export class AfterloginPagesComponent {
  //set user = http get

  //@Input() user: String;

  constructor (private route: Router, private service: UsernameService) {
    
  }

  user:String = this.service.user;

  ngOnInit(): void {}

  navigateToHome() {
    this.route.navigate(['login'])
  }

  navigateToMap() {
    this.route.navigate(['map'])
  }

  navigateToPhotoUpload() {
    this.route.navigate(['photo-upload'])
  }

  navigateToPhotoLib() {
    this.route.navigate(['photo-library'])
  }
}