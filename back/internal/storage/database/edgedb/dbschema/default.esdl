module default {
	abstract type BaseObject {
    required property created_at: datetime {
    	rewrite insert using (datetime_of_statement());
    	readonly := true;
    }
    required property updated_at: datetime {
    	readonly := true;
    	rewrite update,insert using (datetime_of_statement());
    }
    property deleted_at: datetime;
   	}
	type User extending BaseObject {
		required property username: str {
		    constraint exclusive
		};
		required property email: str {
		    constraint exclusive
		};
		required property password: str;
		index on (.username);
	}
}