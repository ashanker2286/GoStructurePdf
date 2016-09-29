RM=rm -f
RMFORCE=rm -rf
SRCS= genGoStructPdf.go
all: exe

exe: $(SRCS)
	go build $(SRCS)

clean:
	$(RM) genGoStructPdf
