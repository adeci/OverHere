import { Component } from '@angular/core';
import { Route, Router } from '@angular/router';

@Component({
  selector: 'app-photo-lib-pages',
  templateUrl: './photo-lib-pages.component.html',
  styleUrls: ['./photo-lib-pages.component.css']
})
export class PhotoLibPagesComponent {
  constructor(private route: Router) {}

  backToHomePage() {
    this.route.navigate(['home'])
  }
}
