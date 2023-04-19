import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { Router, RouterModule, Routes } from '@angular/router';
import { AppRoutingModule } from './app-routing.module';

import { HttpClientModule } from '@angular/common/http';
import { MarkerService } from './marker.service';

import { AppComponent } from './app.component';
import { MapComponent } from './map/map.component';
import { AuthPagesComponent } from './pages/auth-pages/auth-pages.component';
import { LoginPagesComponent } from './pages/login-pages/login-pages.component';
import { NewuserPagesComponent } from './pages/new-user-pages/new-user-pages.component';
import { ReturninguserComponent } from './login-and-signup/returning-user.component';
import { AfterloginComponent } from './home-page/after-login.component';
import { ReturninguserPagesComponent } from './pages/returning-user-pages/returning-user-pages.component';
import { AfterloginPagesComponent } from './pages/after-login-pages/after-login-pages.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatDialogModule } from '@angular/material/dialog';
import {ReactiveFormsModule} from '@angular/forms'
import { NgMultiSelectDropDownModule } from 'ng-multiselect-dropdown';
import { CommonModule } from '@angular/common';
import { PhotoLibComponent } from './photo-lib/photo-lib.component';
import { PhotoLibPagesComponent } from './pages/photo-lib-pages/photo-lib-pages.component';

@NgModule({
  declarations: [
    AppComponent,
    MapComponent,
    AuthPagesComponent,
    LoginPagesComponent,
    NewuserPagesComponent,
    ReturninguserPagesComponent,
    AfterloginPagesComponent,
    PhotoLibPagesComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    RouterModule,
    BrowserAnimationsModule,
    MatDialogModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
    MatFormFieldModule,
    ReactiveFormsModule,
    NgMultiSelectDropDownModule.forRoot(),
    CommonModule
  ],
  exports: [
    MapComponent
  ],
  providers: [
    MarkerService,
    HttpClientModule
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }