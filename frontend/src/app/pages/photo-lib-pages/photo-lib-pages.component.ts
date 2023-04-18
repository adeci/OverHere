import { Component } from '@angular/core';
import { Route, Router } from '@angular/router';
import { HttpClient , HttpErrorResponse} from '@angular/common/http';
import { UsernameService } from 'src/app/username.service';
import { catchError } from 'rxjs';
import { throwError } from 'rxjs';

@Component({
  selector: 'app-photo-lib-pages',
  templateUrl: './photo-lib-pages.component.html',
  styleUrls: ['./photo-lib-pages.component.css']
})
export class PhotoLibPagesComponent {
  constructor(private route: Router, private http:HttpClient, private service:UsernameService) {}

  user:String = this.service.user;
  userid:String = this.service.userid;
  image:string = '';
  index:number = 0;
  //http get user

  private photos:Array<any> = [];
  //http get photos

  backToHomePage() {
    this.route.navigate(['home'])
  }

  getImages() {
    return this.http.get<any>('http://localhost:8000/images/get/byuserid/' + this.userid).pipe(
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

 

  initializeImage() {
    this.getImages().subscribe(
      data => {
        this.photos = data.data.data;
        console.log(this.photos);
        this.image = this.photos[0].encoding;
        this.index = 0;
      },
      error => {
      }
    )
    
  }

  nextImage() {
    this.index++;
    console.log(this.index);
    if (this.index < this.photos.length) {
      this.image = this.photos[this.index].encoding;
    } else {
      return;
    }
  }
}
