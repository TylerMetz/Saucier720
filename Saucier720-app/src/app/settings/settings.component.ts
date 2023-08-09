import { Component } from '@angular/core';
import { ColorService } from '../core/services/color/color.service';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss']
})
export class SettingsComponent{

  constructor(public colorService: ColorService) {}

  setThemeDefault(){
    this.colorService.setColorPalette(ColorService.ColorPalette.Default);
  }
  setThemeCustom1(){
    this.colorService.setColorPalette(ColorService.ColorPalette.Custom1);
  }
  setThemeCustom2(){
    this.colorService.setColorPalette(ColorService.ColorPalette.Custom2);
  }
  setThemeCustom3(){
    this.colorService.setColorPalette(ColorService.ColorPalette.Custom3);
  }
  setThemeCustom4(){
    this.colorService.setColorPalette(ColorService.ColorPalette.Custom4);
  }
}