import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { MapRoutingModule } from './map-routing.module';
import { MapComponent } from './map.component';
import { AppModule } from '../app.module';
import { SharedModule } from '../shared.module';
import { HttpClientModule } from '@angular/common/http';


@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    MapRoutingModule,
    HttpClientModule
  ]
})

export class MapModule {}