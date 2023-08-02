import { Component, Output, EventEmitter, OnInit} from '@angular/core';
import { AuthService } from 'src/app/core/services/Auth/auth.service';
import { trigger, transition, query, style, animate, group } from '@angular/animations';
import { HttpEvent, HttpEventType } from '@angular/common/http';
import { lastValueFrom } from 'rxjs';
import { RecipeService } from 'src/app/core/services/recipes/recipe.service';
import { RecipePost } from 'src/app/core/interfaces/recipe';
import { UserDashboardService } from 'src/app/core/services/user-dashboard/user-dashboard.service';


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

    async ngOnInit() {
      await this.populateDashboardRecipes();
      setTimeout(() => {
        this.generationComplete.emit(true); // event for button generation
      }, 500); // 500ms temporarily
    }

    recipes: RecipePost[] = [];
    currentRecipe!: RecipePost;
    counter: number = 0;
  
    constructor(
      private recipeService: RecipeService,
      private userDashboardService: UserDashboardService,
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
            this.currentRecipe = this.recipes[this.counter];
            break;
        }
      } catch (error) {
        console.error(error);
      }
    }

    calculateRecipeIngredientsOwnedPercentage(recipeIndex: number): number {
      const percentage = (this.recipes[recipeIndex].ItemsInPantry.length / this.recipes[recipeIndex].R.ingredients.length) * 100;
      return Math.round(percentage);
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

    