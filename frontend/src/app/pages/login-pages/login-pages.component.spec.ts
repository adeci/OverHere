import { LoginPagesComponent } from "./login-pages.component";
import { TestBed, async, ComponentFixture, inject, tick, fakeAsync} from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import {RouterTestingModule} from '@angular/router/testing';
import {Router} from '@angular/router';
import {Location} from '@angular/common';

describe('LoginpagesComponent', () => {
  let component : LoginPagesComponent;
  let fixture: ComponentFixture<LoginPagesComponent>;
  let location: Location;
  let router: Router;

    beforeEach(async(() => {
      TestBed.configureTestingModule({
        imports: [ ],
        declarations: [ LoginPagesComponent ],
        providers: [  ]
      }).compileComponents().then(() => {
        fixture = TestBed.createComponent(LoginPagesComponent);
        component = fixture.componentInstance;
      });
    }))

    //base test
    it('should create', () =>  {
      expect(component).toBeTruthy();
    })

    //check if header is printed correctly
    it('should disp title', async(() => {
        const fixture = TestBed.createComponent(LoginPagesComponent);
        fixture.detectChanges();
        const compiled = fixture.debugElement.nativeElement;
        expect(compiled.querySelector('header').textContent).toContain('OverHere');
      }));

    it ('should display buttons',( () => {
      const fixture = TestBed.createComponent(LoginPagesComponent);
        fixture.detectChanges();
        const compiled = fixture.debugElement.nativeElement;
        let list = compiled.querySelectorAll('button');
        expect(list.length).toEqual(2);
        expect(list[0].textContent).toContain("Log In");
        expect(list[1].textContent).toContain("Sign Up");
    }))

    //tests to check if buttons route to correct location
    it ('should link to login on loginbutton click', fakeAsync( () => {
      spyOn(component, 'navigateToReturningUser');
      let list = fixture.debugElement.nativeElement.querySelectorAll('button');
      let button = list[0];
      button.click();

      fixture.whenStable().then(() => {
        expect(component.navigateToReturningUser).toHaveBeenCalled();
      })
    }));

    it ('should link to login on signupbutton click', fakeAsync( () => {
      spyOn(component, 'navigateToNewUser');
      let list = fixture.debugElement.nativeElement.querySelectorAll('button');
      let button = list[1];
      button.click();

      fixture.whenStable().then(() => {
        expect(component.navigateToNewUser).toHaveBeenCalled();
      })
    }));

})