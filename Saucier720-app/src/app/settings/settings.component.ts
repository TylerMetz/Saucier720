import { Component } from '@angular/core';
import { ColorService } from '../core/services/color/color.service';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss']
})
export class SettingsComponent{

  constructor(private colorService: ColorService) {}

  private availablePalettes = [
    ColorService.ColorPalette.Default,
    ColorService.ColorPalette.Custom1,
    ColorService.ColorPalette.Custom2,
    ColorService.ColorPalette.Custom3,
    ColorService.ColorPalette.Custom4,
  ];

  private currentPaletteIndex = 0;
  changeToCustomPalette() {
    this.currentPaletteIndex = (this.currentPaletteIndex + 1) % this.availablePalettes.length;
    const nextPalette = this.availablePalettes[this.currentPaletteIndex];
    this.colorService.setColorPalette(nextPalette);
  }
}