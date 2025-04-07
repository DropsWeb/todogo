PHONY: generate
generate:
	rm -rf pkg/swagger/ops pkg/swagger/models 
		swagger generate server \
			-q \
			--with-enum-ci \
			--exclude-main \
			-f ./swagger-api/swagger.yml \
			-A todo-api \
			-s pkg/swagger \
			-a ops \
			-m pkg/swagger/models \
			--implementation-package github.com/DropsWeb/todogo/internal/services
