import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login-pages',
  templateUrl: './login-pages.component.html',
  styleUrls: ['./login-pages.component.css']
})
export class LoginPagesComponent implements OnInit {
  constructor (private route: Router) {}

  ngOnInit(): void {}

  navigateToNewUser() {
    this.route.navigate(['new-user']);
  }

  navigateToReturningUser() {
    this.route.navigate(['returning-user']);
  }
}
