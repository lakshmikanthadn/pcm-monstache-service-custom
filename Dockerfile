####################################################################################################
# Step 1: Build the app
####################################################################################################

FROM golang 

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go mod download

RUN go get gopkg.in/mgo.v2

RUN go build -buildmode=plugin -o test.so test.go

RUN go build

####################################################################################################
# Step 2: Copy output build file to an alpine image
####################################################################################################


CMD ["./monstache", "-mapper-plugin-path" , "test.so", "-f" ,"config.toml"]

