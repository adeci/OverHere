import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({ selector: 'app-root', templateUrl: 'app.component.html' })
export class AppComponent implements OnInit {
    imgId;
    constructor(private http: HttpClient) { }

    //this is using a test api right now before linking the backend just to demonstrate http requests are being recieved
    ngOnInit() {
      this.http.post<any>('https://reqres.in/api/posts', { title: 'User' }).subscribe(data => {
        this.imgId = data.id;
    })
    }
}

interface Image{
  id: number;
  imgstring: string;
}