import { AfterloginPagesComponent } from './after-login-pages.component';
import { TestBed, async, ComponentFixture, inject, tick, fakeAsync} from '@angular/core/testing';
import { Location } from '@angular/common';
import { Router } from '@angular/router';

describe('AfterloginPagesComponent', () => {
    let component : AfterloginPagesComponent;
    let fixture: ComponentFixture<AfterloginPagesComponent>;
    let location: Location;
    let router: Router;
  
      beforeEach(async(() => {
        TestBed.configureTestingModule({
          imports: [ ],
          declarations: [ AfterloginPagesComponent ],
          providers: [  ]
        }).compileComponents().then(() => {
          fixture = TestBed.createComponent(AfterloginPagesComponent);
          component = fixture.componentInstance;
        });
      }))
  
      //base test
      it('should create afterlogin', () =>  {
        expect(component).toBeTruthy();
      })

      //ensure header is displayed
      it('should disp title', async(() => {
        const fixture = TestBed.createComponent(AfterloginPagesComponent);
        fixture.detectChanges();
        const compiled = fixture.debugElement.nativeElement;
        expect(compiled.querySelector('header').textContent).toContain('Welcome,');
      }));


      //ensure all buttons link correctly
      it ('should link to home on signout click', fakeAsync( () => {
        spyOn(component, 'navigateToHome');
        let list = fixture.debugElement.nativeElement.querySelectorAll('button');
        let button = list[0];
        button.click();
  
        fixture.whenStable().then(() => {
          expect(component.navigateToHome).toHaveBeenCalled();
        })
      }));

      it ('should link to map on mapbutton click', fakeAsync( () => {
        spyOn(component, 'navigateToMap');
        let list = fixture.debugElement.nativeElement.querySelectorAll('button');
        let button = list[1];
        button.click();
  
        fixture.whenStable().then(() => {
          expect(component.navigateToMap).toHaveBeenCalled();
        })
      }));

      it ('should link to photoup on photoup click', fakeAsync( () => {
        spyOn(component, 'navigateToPhotoUpload');
        let list = fixture.debugElement.nativeElement.querySelectorAll('button');
        let button = list[2];
        button.click();
  
        fixture.whenStable().then(() => {
          expect(component.navigateToPhotoUpload).toHaveBeenCalled();
        })
      }));

      it ('should link to lib on library click', fakeAsync( () => {
        spyOn(component, 'navigateToPhotoLib');
        let list = fixture.debugElement.nativeElement.querySelectorAll('button');
        let button = list[3];
        button.click();
  
        fixture.whenStable().then(() => {
          expect(component.navigateToPhotoLib).toHaveBeenCalled();
        })
      }));


})