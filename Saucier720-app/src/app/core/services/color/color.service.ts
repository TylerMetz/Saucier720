import { Injectable } from '@angular/core';


@Injectable({
    providedIn: 'root',
  })
  export class ColorService {
    static ColorPalette = {
      Default: 'default',
      Custom1: 'custom1',
      Custom2: 'custom2',
    };
  
    private colorPalettes = {
        [ColorService.ColorPalette.Default]: {
            '--primary-color': '#141451',
            '--secondary-color': '#ec8091',
            '--text-color': '#f8f0ca',
            '--background-color': '#0e3c2c',
            '--header-color': '#549464',
            '--primary-color-image': "url('/assets/images/--primary-color-default.png')",
            '--secondary-color-image': "url('/assets/images/--secondary-color-default.png')",
            '--text-color-image': "url('/assets/images/--text-color-default.png')",
            '--background-color-image': "url('/assets/images/--background-color-default.png')",
            '--header-color-image': "url('/assets/images/--header-color-default.png')",
        },
        [ColorService.ColorPalette.Custom1]: {
            '--primary-color': '#033043',
            '--secondary-color': '#E9E3D5',
            '--text-color': '#FDA521',
            '--background-color': '#719D99',
            '--header-color': '#0A7273',
            '--primary-color-image': "url('/assets/images/--primary-color-custom1.png')",
            '--secondary-color-image': "url('/assets/images/--secondary-color-custom1.png')",
            '--text-color-image': "url('/assets/images/--text-color-custom1.png')",
            '--background-color-image': "url('/assets/images/--background-color-custom1.png')",
            '--header-color-image': "url('/assets/images/--header-color-custom1.png')",
        },
    };
    
    // ... other methods
  
    setColorPalette(palette: string) {
      // ... existing implementation
    }
  }
  