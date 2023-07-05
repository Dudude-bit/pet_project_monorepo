module default {
	abstract type BaseObject {
    required property created_at: cal::local_datetime {
    	rewrite insert using (datetime_of_statement());
    	readonly := true;
    }
    required property updated_at: cal::local_datetime {
    	readonly := true;
    	rewrite update using (datetime_of_statement());
    }
    property deleted_at: cal::local_datetime;
   	}
	type User extending BaseObject {
		required property username: str;
		required property email: str;
		required property password: str;
	}
}
