import { PhotoLibPagesComponent } from "./photo-lib-pages.component";
import { TestBed, async, ComponentFixture, inject, tick, fakeAsync} from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import {RouterTestingModule} from '@angular/router/testing';
import {Router} from '@angular/router';
import {Location} from '@angular/common'
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { HttpClientModule } from '@angular/common/http';

describe('PhotoLibPagesComponent', () => {
    let component : PhotoLibPagesComponent;
    let fixture: ComponentFixture<PhotoLibPagesComponent>;
    let location: Location;
    let router: Router;

    beforeEach(async(() => {
      TestBed.configureTestingModule({
        imports: [HttpClientModule],
        declarations: [ PhotoLibPagesComponent ],
        providers: [  ]
      }).compileComponents().then(() => {
        fixture = TestBed.createComponent(PhotoLibPagesComponent);
        component = fixture.componentInstance;
      });
    }))

    //base test
    it('should create photolib', () =>  {
      expect(component).toBeTruthy();
    })

    //check if header is printed correctly
    it('should disp photolib title', async(() => {
        const fixture = TestBed.createComponent(PhotoLibPagesComponent);
        fixture.detectChanges();
        const compiled = fixture.debugElement.nativeElement;
        expect(compiled.querySelector('header').textContent).toContain('Your OverHere Photo Library');
      }));

    it ('should display nextimg and back buttons',( () => {
        const fixture = TestBed.createComponent(PhotoLibPagesComponent);
        fixture.detectChanges();
        const compiled = fixture.debugElement.nativeElement;
        let list = compiled.querySelectorAll('button');
        expect(list.length).toEqual(2)
        expect(list[0].textContent).toContain("Next Image");
        expect(list[1].textContent).toContain("Back");
    }))

    it ('should link to afterlogin on back click', fakeAsync( () => {
      spyOn(component, 'backToHomePage');
      let list = fixture.debugElement.nativeElement.querySelectorAll('button');
      let button = list[1];
      button.click();

      fixture.whenStable().then(() => {
        expect(component.backToHomePage).toHaveBeenCalled();
      })
    }));
})

