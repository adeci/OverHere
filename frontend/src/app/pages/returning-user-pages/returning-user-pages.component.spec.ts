import { ReturninguserPagesComponent } from './returning-user-pages.component';
import { TestBed, async, ComponentFixture, inject, tick, fakeAsync} from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http';
import { Router } from '@angular/router';

describe('ReturninguserPagesComponent', () => {
    let component : ReturninguserPagesComponent;
    let fixture: ComponentFixture<ReturninguserPagesComponent>;
    let location: Location;
    let router: Router;

    beforeEach(async(() => {
        TestBed.configureTestingModule({
          imports: [HttpClientModule],
          declarations: [ ReturninguserPagesComponent ],
          providers: [  ]
        }).compileComponents().then(() => {
          fixture = TestBed.createComponent(ReturninguserPagesComponent);
          component = fixture.componentInstance;
        });
      }))

    it('should create returninguser', () =>  {
        expect(component).toBeTruthy();
    })

    it('should disp returning title', async(() => {
        const fixture = TestBed.createComponent(ReturninguserPagesComponent);
        fixture.detectChanges();
        const compiled = fixture.debugElement.nativeElement;
        expect(compiled.querySelector('header').textContent).toContain('Welcome Back!');
      }));

    it ('should display submit and back buttons',( () => {
        const fixture = TestBed.createComponent(ReturninguserPagesComponent);
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