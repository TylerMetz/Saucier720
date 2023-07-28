import { Component, ElementRef, ViewChild, Renderer2 } from '@angular/core';
import { trigger, transition, style, animate } from '@angular/animations';

@Component({
    selector: 'app-landing-page',
    templateUrl: './landing-page.component.html',
    styleUrls: ['./landing-page.component.scss']
  })

  export class LandingPageComponent{
    buttonsVisible: boolean = false;
    isComponentReady = false;


    onLogoGenerationComplete(generationComplete: boolean) {
      if (generationComplete) {
        // The logo generation is complete, you can now trigger the buttons generation or their appearance.
        // For example, you can set a flag to control their visibility:
        this.buttonsVisible = true;
        this.isComponentReady = true;
      }
    }

  activeButton: string = '';
  hoveredButton: string = '';

  isButtonClicked(buttonName: string): boolean {
    return this.activeButton === buttonName;
  }

  onButtonClick(buttonName: string): void {
    this.activeButton = this.activeButton === buttonName ? '' : buttonName;
  }

  onButtonHover(buttonName: string): void {
    this.hoveredButton = buttonName;
  }

  onButtonLeave(buttonName: string): void {
    this.hoveredButton = '';
  }

  isButtonHovered(buttonName: string): boolean {
    return this.hoveredButton === buttonName;
  }

  }
  