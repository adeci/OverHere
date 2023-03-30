import { NewuserPagesComponent } from "./new-user-pages.component";
import { TestBed, async, ComponentFixture, inject, tick, fakeAsync} from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import {RouterTestingModule} from '@angular/router/testing';
import {Router} from '@angular/router';
import {Location} from '@angular/common'
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { HttpClientModule } from '@angular/common/http';

describe('NewuserPagesComponent', () => {
    let component : NewuserPagesComponent;
    let fixture: ComponentFixture<NewuserPagesComponent>;
    let location: Location;
    let router: Router;

    beforeEach(async(() => {
      TestBed.configureTestingModule({
        imports: [HttpClientModule],
        declarations: [ NewuserPagesComponent ],
        providers: [  ]
      }).compileComponents().then(() => {
        fixture = TestBed.createComponent(NewuserPagesComponent);
        component = fixture.componentInstance;
      });
    }))

    //base test
    it('should create newuser', () =>  {
      expect(component).toBeTruthy();
    })

    //check if header is printed correctly
    it('should disp new title', async(() => {
        const fixture = TestBed.createComponent(NewuserPagesComponent);
        fixture.detectChanges();
        const compiled = fixture.debugElement.nativeElement;
        expect(compiled.querySelector('header').textContent).toContain('OverHere');
      }));

    it ('should display submit and back buttons',( () => {
        const fixture = TestBed.createComponent(NewuserPagesComponent);
        fixture.detectChanges();
        const compiled = fixture.debugElement.nativeElement;
        let list = compiled.querySelectorAll('button');
        expect(list.length).toEqual(2)
        expect(list[0].textContent).toContain("Confirm");
        expect(list[1].textContent).toContain("Back");
    }))

    it ('should link to afterlogin on confirm click', fakeAsync( () => {
      spyOn(component, 'navigateToHomePage');
      let list = fixture.debugElement.nativeElement.querySelectorAll('button');
      let button = list[0];
      button.click();

      fixture.whenStable().then(() => {
        expect(component.navigateToHomePage).toHaveBeenCalled();
      })
    }));

    it ('should link to login on back click', fakeAsync( () => {
      spyOn(component, 'navigateToHome');
      let list = fixture.debugElement.nativeElement.querySelectorAll('button');
      let button = list[1];
      button.click();

      fixture.whenStable().then(() => {
        expect(component.navigateToHome).toHaveBeenCalled();
      })
    }));

})