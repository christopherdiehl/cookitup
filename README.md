# CookItUp

CookItUp is a cli tool to store your recipes and generate a random meal plan for the week. 

Vision:
Services like Blue Apron and Hello Fresh provide the convenience of a meal plan and food for a price. This tool aims to provide the service of a meal plan and grocery list without the associated cost of Hello Fresh or Blue Apron. This is inspired by my wife and I wanting not wanting to create a grocery list for meals and relying on Hello Fresh instead, even though we still need to shop for the remaining meals.

Goals:
- Able to store recipes in JSON format
- Specify number of meals in random meal plan
- Randomly generate meal plan 
- Send 2 text messages to your specified phone number using twilio upon meal plan generation (optional via flag)
    1.  The ingredients of the recipes to your phone (useful for grocery lists).
    2.  The full recipes for the generated mealplan (useful for cooking).
