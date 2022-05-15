## Database design:

<div align="center">
  <img src="https://github.com/raflynagachi/crowdfunding-web/blob/master/assets/erd.png"/><br>
  <caption><i>Entity Relationship Diagram</i></caption>
</div><br>

Attributes:
- Users
  - id int
  - name varchar
  - email varchar
  - password_hash
  - occupation varchar
  - avatar_filename varchar
  - role varchar
  - remember_token varchar
  - created_at datetime
  - updated_at datetime
- Campaigns
  - id int
  - user_id int
  - name varchar
  - slug varchar
  - short_description varchar
  - description text
  - goal_amount float(decimal)
  - current_amount float(decimal)
  - backer_count int
  - perks text
  - created_at datetime
  - updated_at datetime
- Campaign Images
  - id int
  - campaign_id int
  - filename varchar
  - is_primary tinyint(boolean)
  - created_at datetime
  - updated_at datetime
- Transactions
  - id int
  - user_id int
  - campaign_id int
  - amount float
  - status varchar
  - code varchar
  - created_at datetime
  - updated_at datetime