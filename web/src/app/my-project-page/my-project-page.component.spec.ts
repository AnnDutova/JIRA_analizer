import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MyProjectPageComponent } from './my-project-page.component';

describe('MyProjectPageComponent', () => {
  let component: MyProjectPageComponent;
  let fixture: ComponentFixture<MyProjectPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MyProjectPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MyProjectPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
