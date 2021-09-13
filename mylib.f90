module hi

contains
    subroutine hello() bind(C)
        print *, "Hello from Fortran"
    end subroutine

end module hi
