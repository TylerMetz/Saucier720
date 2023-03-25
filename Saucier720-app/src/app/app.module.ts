import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppComponent } from './app.component';
import { AppRoutingModule, routingComponents } from './app-routing.module';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { DealsTableComponent } from './deals/FEC/deals-table/deals-table.component';
import { PantryTableComponent } from './pantry/FEC/pantry-table/pantry-table.component';

@NgModule({
  declarations: [
    AppComponent,
    routingComponents,
    LoginComponent,
    SignupComponent,
    DealsTableComponent,
    PantryTableComponent
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
