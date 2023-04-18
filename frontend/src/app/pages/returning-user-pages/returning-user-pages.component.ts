import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RouterModule } from '@angular/router';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { UsernameService } from 'src/app/username.service';
import { catchError } from 'rxjs';
import { throwError } from 'rxjs';

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
  show = false;
  private req:string = "";
  isValid:Boolean = false;

  ngOnInit(): void {}

  navigateToHome() {
    this.currentuser=''
    this.route.navigate(['login'])
  }

  getUser() {
    return this.http.get<any>('http://localhost:8000/users/get/byusername/' + this.currentuser).pipe(
      catchError((error: HttpErrorResponse) => {
        let errorMessage = 'Unknown error occurred';
        if (error.error instanceof ErrorEvent) {
          // Client-side error
          errorMessage = `Error: ${error.error.message}`;
        } else {
          // Server-side error
          errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
        }
        return throwError(errorMessage);
      })
    );
  }
  
  navigateToHomePage(val:string) {
    this.currentuser=val
    this.service.user = this.currentuser;
    this.getUser().subscribe(
      data => {
        this.req = data.data.data.username;
        this.service.userid = data.data.data.userid;
        this.route.navigate(['home'], {state: {data:val}})
        return;
      },
      error => {
        console.log("error");
        this.show = true;
        this.req = '';
        return;
      }
    )
  }

  closeErrorPopup() {
    this.show = false;
  }

  navigateToNewUser() {
    this.route.navigate(['new-user']);
  }
}