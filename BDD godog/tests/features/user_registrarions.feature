Feature: user registration
    In this feature we will handle user registration

    Scenario: user registration valid request
        When I try to register user with age 20 and name: Nick
        Then I expect not to see an error 

    Scenario: user registration with not valid age
        When I try to register user with age 15 and name: Nick
        Then I expect not to see an error 