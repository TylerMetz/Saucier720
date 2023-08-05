import { Component, Output, EventEmitter, OnInit, Renderer2, ElementRef, ViewChild} from '@angular/core';
import { AuthService } from 'src/app/core/services/Auth/auth.service';
import { trigger, transition, query, style, animate, group } from '@angular/animations';
import { HttpEvent, HttpEventType } from '@angular/common/http';
import { lastValueFrom } from 'rxjs';
import { RecipeService } from 'src/app/core/services/recipes/recipe.service';
import { RecipePost } from 'src/app/core/interfaces/recipe';
import { UserDashboardService } from 'src/app/core/services/user-dashboard/user-dashboard.service';
import { ImagesService } from 'src/app/core/services/images/images.service';

const left = [
  query(':enter, :leave', style({ position: 'absolute'}), { optional: true }),
  group([
    query(':enter', [style({ transform: 'translateX(-150%)' }), animate('.5s ease-out', style({ transform: 'translateX(0%)' }))], {
      optional: true,
    }),
    query(':leave', [style({ transform: 'translateX(0%)' }), animate('.5s ease-out', style({ transform: 'translateX(150%)' }))], {
      optional: true,
    }),
  ]),
];

const right = [
  query(':enter, :leave', style({ position: 'absolute' }), { optional: true }),
  group([
    query(':enter', [style({ transform: 'translateX(150%)' }), animate('.5s ease-out', style({ transform: 'translateX(0%)' }))], {
      optional: true,
    }),
    query(':leave', [style({ transform: 'translateX(0%)' }), animate('.5s ease-out', style({ transform: 'translateX(-150%)' }))], {
      optional: true,
    }),
  ]),
];

@Component({
    selector: 'app-user-dashboard',
    templateUrl: './user-dashboard.component.html',
    styleUrls:[ './user-dashboard.component.scss'],
    animations: [
      trigger('animSlider', [
        transition('0 => 1', right),
        transition('1 => 0', left),
        transition('* => 0', right), // wrapping from last entry to first
        transition('0 => *', left), // wrapping from first entry to last
        transition(':increment', right),
        transition(':decrement', left),
        
      ]),
    ],
  })

  export class UserDashboardComponent implements OnInit{
    
    // used for button generation
    @Output() generationComplete = new EventEmitter<boolean>();

    // used for button repositioning
    @Output() cardExpanded = new EventEmitter<boolean>();

    // button
    @ViewChild('recipeCardButton') recipeCardButton!: ElementRef;
    
    async ngOnInit() {
      await this.populateDashboardRecipes();
      setTimeout(() => {
        this.generationComplete.emit(true); // event for button generation
      }, 500); // 500ms temporarily
    }

    showTitle: boolean = false;
    recipes: RecipePost[] = [];
    currentRecipe!: RecipePost;
    counter: number = 0;
    hasError: boolean = false;
    showRecipe: boolean = false;
  
    constructor(
      private recipeService: RecipeService,
      private userDashboardService: UserDashboardService,
      private imagesService: ImagesService,
      private renderer: Renderer2, 
      private elementRef: ElementRef
    ) {}
  
    public async populateDashboardRecipes(): Promise<void> {
      try {
        const event: HttpEvent<any> = await lastValueFrom(this.userDashboardService.getUserDashboardData());
        switch(event.type) {
          case HttpEventType.Sent:
            console.log('Request sent!');
            break;
          case HttpEventType.ResponseHeader:
            console.log('Response header received!');
            break;
          case HttpEventType.DownloadProgress:
            const kbLoaded = Math.round(event.loaded / 1024);
            console.log(`Download in progress! ${kbLoaded}Kb loaded`);
            break;
          case HttpEventType.Response:
            console.log('Done!', event.body);
            let recipeStr = JSON.stringify(event.body);
            let parsedRecipes = JSON.parse(recipeStr);
            this.recipes = parsedRecipes;
            // check if recipes are there
            if (!this.recipes) {
              this.hasError = true; // Set the error flag
            } else {
              this.hasError = false; // Clear the error flag
              this.currentRecipe = this.recipes[this.counter];
            }
            break;
        }
      } catch (error) {
        console.error(error);
      }
    }

    calculateRecipeIngredientsOwnedPercentage(recipeIndex: number): number {
      if (this.recipes[recipeIndex].ItemsInPantry.length === 0 || this.recipes[recipeIndex].R.ingredients.length === 0) {
        return 0;
      }
      const percentage = (this.recipes[recipeIndex].ItemsInPantry.length / this.recipes[recipeIndex].R.ingredients.length) * 100;
      return Math.round(percentage);
    }

    searchImage(recipeTitle: string): string{
      let imageUrl = '';
      this.imagesService.searchImage(recipeTitle).subscribe((data: any) => {
        if (data.photos && data.photos.length > 0) {
          imageUrl = data.photos[0].src.large;
        }
      });
      return imageUrl;
    }

    getImages(recipes: RecipePost[]): void {
      recipes.forEach(recipe => {
        recipe.R.pictureLink = this.searchImage(recipe.R.title);
      });
    }

    sizeRecipePreviewWindow() {
      const recipeCardPreview = this.elementRef.nativeElement.querySelector('.highlighted-recipes-box');
      
      if (recipeCardPreview) {
        if (recipeCardPreview.classList.contains('expanded')) {
          recipeCardPreview.classList.remove('expanded');
          this.recipeCardButton.nativeElement.innerText = 'Expand Recipe';
          this.showRecipe = false;
          this.cardExpanded.emit(false);
        } else {
          recipeCardPreview.classList.add('expanded');
          this.recipeCardButton.nativeElement.innerText = 'Collapse Recipe';
          this.showRecipe = true;
          this.cardExpanded.emit(true);
        }
      }
    }
    
    public getAuthorCreditFromRecipeID(recipeID: string): string {
      // used to get recipe author from recipeID
      const author = recipeID.replace(/\d+/g, '');
      if (author === 'json') {
        return 'MealDealz Classic Recipe';
      } else {
        return 'Created by ' + author;
      }
    }

    removeQuotesAndBrackets(arr: string[]): string[] {
      const regex = /["\[\]]/g; // Matches any occurrence of ", [, or ] globally
      return arr.map(str => str.replace(regex, '')); // Replace all matches in each string in the array
    }

    // animation functions 
    onNext() {
      if (this.counter != this.recipes.length - 1) {
        this.counter++;
      } else{
        this.counter = 0;
      }      
    }
    onPrevious() {
      if (this.counter > 0) {
        this.counter--;
      } else{
        this.counter = this.recipes.length - 1;
      }
    }

    
    
}

    