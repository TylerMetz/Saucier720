import { Injectable } from '@angular/core';


@Injectable({
    providedIn: 'root',
  })
  export class ColorService {
    static ColorPalette = {
      Default: 'default',
      Custom1: 'custom1',
      Custom2: 'custom2',
    Custom3: 'custom3',
    Custom4: 'custom4',
    };
  
    private colorPalettes = {
        [ColorService.ColorPalette.Default]: {
            '--primary-color': '#141451',
            '--secondary-color': '#f8f0ca',
            '--text-color': '#ec8091',
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
        [ColorService.ColorPalette.Custom2]: {
            '--primary-color': '#05161A',
            '--secondary-color': '#cbcbcb',
            '--text-color': '#6DA5C0',
            '--background-color': '#0C7075',
            '--header-color': '#0F969C',
            '--primary-color-image': "url('/assets/images/--primary-color-custom2.png')",
            '--secondary-color-image': "url('/assets/images/--secondary-color-custom2.png')",
            '--text-color-image': "url('/assets/images/--text-color-custom2.png')",
            '--background-color-image': "url('/assets/images/--background-color-custom2.png')",
            '--header-color-image': "url('/assets/images/--header-color-custom2.png')",
        },
        [ColorService.ColorPalette.Custom3]: {
            '--primary-color': '#190019',
            '--secondary-color': '#FBE4D8',
            '--text-color': '#DFB6B2',
            '--background-color': '#854F6C',
            '--header-color': '#522B5B',
            '--primary-color-image': "url('/assets/images/--primary-color-custom3.png')",
            '--secondary-color-image': "url('/assets/images/--secondary-color-custom3.png')",
            '--text-color-image': "url('/assets/images/--text-color-custom3.png')",
            '--background-color-image': "url('/assets/images/--background-color-custom3.png')",
            '--header-color-image': "url('/assets/images/--header-color-custom3.png')",
        },
        [ColorService.ColorPalette.Custom4]: {
            '--primary-color': '#161E2F',
            '--secondary-color': '#FFA586',
            '--text-color': '#4a9dea',
            '--background-color': '#541A2E',
            '--header-color': '#851A2B',
            '--primary-color-image': "url('/assets/images/--primary-color-custom4.png')",
            '--secondary-color-image': "url('/assets/images/--secondary-color-custom4.png')",
            '--text-color-image': "url('/assets/images/--text-color-custom4.png')",
            '--background-color-image': "url('/assets/images/--background-color-custom4.png')",
            '--header-color-image': "url('/assets/images/--header-color-custom4.png')",
        },
    };
    
    // ... other methods
  
    setColorPalette(palette: string) {
        const selectedPalette = this.colorPalettes[palette];
    
        if (selectedPalette) {
          Object.entries(selectedPalette).forEach(([property, value]) => {
            document.documentElement.style.setProperty(property, value);
          });
        }
      }
  }
  