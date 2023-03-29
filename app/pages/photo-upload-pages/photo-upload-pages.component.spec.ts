import { PhotoUpPagesComponent } from './photo-upload-pages.component';
import { TestBed, async, ComponentFixture, inject, tick, fakeAsync} from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http';
import { Router } from '@angular/router';
import { Location } from '@angular/common';

describe('PhotoupPagesComponent', () => {
  let component : PhotoUpPagesComponent;
  let fixture: ComponentFixture<PhotoUpPagesComponent>;
  let location: Location;
  let router: Router;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientModule],
      declarations: [ PhotoUpPagesComponent ],
      providers: [  ]
    }).compileComponents().then(() => {
      fixture = TestBed.createComponent(PhotoUpPagesComponent);
      component = fixture.componentInstance;
    });
  }))

    it('should create photoup', () =>  {
        expect(component).toBeTruthy();
    })

    it ('should link to homepage on back click', fakeAsync( () => {
      spyOn(component, 'navigateToHomePage');
      let list = fixture.debugElement.nativeElement.querySelectorAll('button');
      let button = list[1];
      button.click();

      fixture.whenStable().then(() => {
        expect(component.navigateToHomePage).toHaveBeenCalled();
      })
    }));
})