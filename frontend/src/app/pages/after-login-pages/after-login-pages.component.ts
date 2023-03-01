import { Component, Input } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-after-login-pages',
  templateUrl: './after-login-pages.component.html',
  styleUrls: ['./after-login-pages.component.css']
})
export class AfterloginPagesComponent {
  user:string = history.state.data

  //@Input() user: String;

  constructor (private route: Router) {}

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